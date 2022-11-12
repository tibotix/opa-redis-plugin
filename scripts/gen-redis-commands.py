#!/usr/bin/env python3

import enum
import sys
import re
from pathlib import Path
from typing import List

commands_file = Path(__file__).parent.parent / "vendor" / "github.com" / "go-redis" / "redis" / "v8" / "commands.go"
cmdable_class_regex = r"type Cmdable interface {((.|\s)*?)\s}"
skip_command_names = (
    "Pipeline",
    "Pipelined",
    "TxPipelined",
    "TxPipeline",
    "Command",
    "Scan",
    "ScanType",
    "SScan",
    "HScan",
    "ZScan",
    "Eval",
    "EvalSha"
)
skip_go_types = (
    "*Sort",
    "*XPendingExtArgs",
    "*ZStore",
    "ZStore",
    "LPosArgs",
    "*Z",
    "ZAddArgs",
    "ZRangeArgs",
    "*ZRangeBy",
    "*XAddArgs",
    "SetArgs",
    "*BitCount",
    "*XClaimArgs",
    "*GeoRadiusQuery",
    "*GeoSearchQuery",
    "*GeoSearchLocationQuery",
    "*GeoSearchStoreQuery"
)
skip_cmd_classes = (
    "SliceCmd",
    "StringStringMapCmd",
    "StringIntMapCmd",
    "StringStructMapCmd",
    "XMessageSliceCmd",
    "XStreamSliceCmd",
    "XPendingCmd",
    "XPendingExtCmd",
    "XAutoClaimCmd",
    "XAutoClaimJustIDCmd",
    "XInfoGroupsCmd",
    "XInfoStreamCmd",
    "XInfoStreamFullCmd",
    "XInfoConsumersCmd",
    "ZWithKeyCmd",
    "ZSliceCmd",
    "TimeCmd",
    "ClusterSlotsCmd",
    "GeoPosCmd",
    "GeoLocationCmd",
    "GeoSearchLocationCmd"
)


class RegoScalarType(enum.Enum):
    STRING = "String"
    INTNUMBER = "IntNumber"
    UINTNUMBER = "UIntNumber"
    FLOATNUMBER = "FloatNumber"
    BOOL = "Bool"

    @classmethod
    def from_go_type(cls, go_type: str):
        mapping = {
            "bool": cls.BOOL,
            "string": cls.STRING,
            "int": cls.INTNUMBER,
            "int64": cls.INTNUMBER,
            "uint64": cls.UINTNUMBER,
            "float64": cls.FLOATNUMBER,
            "interface{}": cls.STRING,
            "time.Duration": cls.INTNUMBER,
            "time.Time": cls.INTNUMBER,
        }
        return mapping[go_type]
    
    @classmethod
    def from_redis_cmd_class(cls, cmd_class):
        mapping = {
            "StringCmd": cls.STRING,
            "IntCmd": cls.INTNUMBER,
            "BoolCmd": cls.BOOL,
            "DurationCmd": cls.INTNUMBER,
            "StatusCmd": cls.STRING,
            "FloatCmd": cls.FLOATNUMBER,

        }
        return mapping.get(cmd_class, None)
    
    def to_go_rego_types_api_code(self):
        mapping = {
            self.STRING.value: "String",
            self.INTNUMBER.value: "Number",
            self.UINTNUMBER.value: "Number",
            self.FLOATNUMBER.value: "Number",
            self.BOOL.value: "Boolean"
        }
        return f"types.{mapping[self.value]}{{}}"
    
    def to_go_ast_term(self, var_name):
        mapping = {
            self.STRING.value: "StringTerm({var_name})",
            self.BOOL.value: "BooleanTerm({var_name})",
            self.INTNUMBER.value: "IntNumberTerm(int({var_name}))",
            self.UINTNUMBER.value: "UintNumberTerm(uint({var_name}))",
            self.FLOATNUMBER.value: "FloatNumberTerm(float64({var_name}))"
        }
        return f"ast.{mapping[self.value].format(var_name=var_name)}"

class RegoType:
    @classmethod
    def from_go_type(cls, go_type):
        is_composite = "..." in go_type
        go_type = go_type.replace("...", "")
        scalar_type = RegoScalarType.from_go_type(go_type)
        return cls(go_type, scalar_type, is_composite)
    
    @classmethod
    def from_redis_cmd_class(cls, cmd_class):
        cmd_class = cmd_class.replace("*", "").strip()
        skip = cmd_class in skip_cmd_classes
        is_composite = "Slice" in cmd_class
        cmd_class = cmd_class.replace("Slice", "").strip()
        scalar_type = RegoScalarType.from_redis_cmd_class(cmd_class)
        return cls(cmd_class, scalar_type, is_composite), skip
        
    def __init__(self, go_type: str, scalar_type: RegoScalarType, is_composite: bool):
        self.go_type = go_type
        self.scalar_type = scalar_type
        self.is_composite = is_composite
    
    def to_go_rego_types_api_code(self):
        if not self.is_composite:
            return self.scalar_type.to_go_rego_types_api_code()
        return f"types.NewArray([]types.Type{{}}, {self.scalar_type.to_go_rego_types_api_code()})"
    
    def to_go_var_declaration(self, var_name):
        if not self.is_composite:
            return f"var {var_name} {self.go_type}"
        return f"var {var_name} []{self.go_type}"
    
    def to_go_parameter_code(self, var_name):
        if not self.is_composite or "interface" in self.go_type:
            return var_name
        return f"{var_name}..."
    
    def to_go_ast_term_return_statement(self, var_name):
        if not self.is_composite:
            return f"return {self.scalar_type.to_go_ast_term(var_name)}, nil"
        return f"""
        var ret []*ast.Term
        for _, v := range {var_name} {{
            ret = append(ret, {self.scalar_type.to_go_ast_term("v")})
        }}
        return ast.ArrayTerm(ret...), nil
                """

class CmdSignature:
    @classmethod
    def from_string(cls, signature):
        # print(f"Parsing Signature: {signature!s}")
        cmd_name = re.match(r"(^.*?)\(", signature).group(1)
        signature = signature[len(cmd_name)+1:].strip()
        signature = re.sub(r"ctx context\.Context(, )?", "", signature)
        parameter_types, skip_go_type = CmdSignature.parse_parameter_list(signature[:signature.index(")")].split(","))
        signature = signature[signature.index(")")+1:].strip()
        return_type, skip_cmd_class = RegoType.from_redis_cmd_class(signature)

        # print(f"cmd_name: {cmd_name!s}")
        # print(f"parameter_types: {parameter_types!s}")
        # print(f"return_type: {signature!s}")
        return cls(cmd_name, parameter_types, return_type, skip_go_type or skip_cmd_class)
    
    @staticmethod
    def parse_parameter_list(parameters: List[str]):
        skip = False
        parameter_types = list()
        parameters = list(filter(lambda x:x, map(lambda x:x.strip(), reversed(parameters))))
        for p in parameters:
            parts = p.split(" ")
            if len(parts) <= 1:
                parameter_types.insert(0, parameter_types[0])
                continue
            go_type = parts[-1]
            if go_type in skip_go_types:
                skip = True
            try:
                parameter_types.insert(0, RegoType.from_go_type(go_type))
            except KeyError as e:
                print(f"KEY ERROR: {e!s}", file=sys.stderr)
                continue
        return parameter_types, skip
            

    def __init__(self, cmd_name: str, parameter_types: List[RegoType], return_type: RegoType, skip: bool):
        self.cmd_name = cmd_name
        self.parameter_types = parameter_types
        self.return_type = return_type
        self.skip = skip


def main():
    if not commands_file.exists():
        print(f"File: {commands_file.absolute()} not found", file=sys.stderr)
        return 1
    if (cmdable := re.search(cmdable_class_regex, commands_file.read_text())) is None:
        print("Could not find Cmdable interface", file=sys.stderr)
        return 1
    command_signatures = cmdable.group(1).split("\n")
    
    print("// Code generated! DO NOT EDIT")
    print("package internal")
    print("""
import (
    "time"

    "github.com/go-redis/redis/v8"
    "github.com/open-policy-agent/opa/ast"
    "github.com/open-policy-agent/opa/rego"
    "github.com/open-policy-agent/opa/types"
)
    """)
    register_functions = list()
    for signature in filter(lambda s:s.cmd_name not in skip_command_names, parse_command_signatures(command_signatures)):
        if signature.skip: continue
        # print(f" Generating Function Code for: {signature.cmd_name}")
        code = generate_function_code_for_signature(signature)
        register_functions.append(f"register{signature.cmd_name.upper()}")
        print(code)

    register_main_func = """
func (p *redisPlugin) registerRedisCommands() {"""
    for f in register_functions:
        register_main_func += f"""
    {f}(p)"""
    register_main_func += """
}"""
    print(register_main_func)

    return 0


def parse_command_signatures(command_signatures):
    for signature in map(lambda x:x.strip(), filter(lambda x:x and "//" not in x, command_signatures)):
        yield CmdSignature.from_string(signature)


def generate_function_code_for_signature(signature: CmdSignature):
    args = ",".join(map(lambda p:p.to_go_rego_types_api_code(), signature.parameter_types))

    code = f"""
func register{signature.cmd_name.upper()}(p *redisPlugin) {{
    rego.RegisterBuiltinDyn(
        &rego.Function{{
            Name: \"redis.{signature.cmd_name.lower()}\",
            Decl: types.NewFunction(types.Args({args}), {signature.return_type.to_go_rego_types_api_code()}),
            Memoize: true,
            Nondeterministic: true,
        }},
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {{
            rdb, err := p.redisProxy.Get()
            if err != nil {{
                return nil, err
            }}

"""

    parameter_args = ["p.redisContext"]
    for idx, p in enumerate(signature.parameter_types):
        var_name = f"v{idx!s}"
        parameter_args.append(p.to_go_parameter_code(var_name))
        code += f"""
            {p.to_go_var_declaration(var_name)}
            if err := ast.As(terms[{idx!s}].Value, &{var_name}); err != nil {{
                return nil, err
            }}
            """
    parameter_args = ",".join(parameter_args)

    code += f"""

            val, err := rdb.{signature.cmd_name}({parameter_args}).Result()
            switch {{
            case err == redis.Nil:
                return ast.NullTerm(), nil
            case err != nil:
                return nil, err
            default:
                {signature.return_type.to_go_ast_term_return_statement("val")}
            }}
        }},
    )
}}
"""
    return code


if __name__ == "__main__":
    sys.exit(main())
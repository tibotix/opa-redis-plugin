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
    "ClusterSlotsCmd",
    "GeoPosCmd",
    "GeoLocationCmd",
    "GeoSearchLocationCmd"
)
# TODO: implement some of these


class RegoScalarType(enum.Enum):
    ANY = 0
    STRING = 1
    INTNUMBER = 2
    UINTNUMBER = 3
    FLOATNUMBER = 4
    BOOL = 5
    DURATION = 6
    TIME = 7

class RegoTypeConvertible:
    def __init__(self, go_type, scalar_type, composite):
        self.go_type = go_type
        self.scalar_type = scalar_type
        self.is_composite = composite

    def _to_go_rego_types_api_code(self):
        mapping = {
            RegoScalarType.ANY: "Any",
            RegoScalarType.STRING: "String",
            RegoScalarType.INTNUMBER: "Number",
            RegoScalarType.UINTNUMBER: "Number",
            RegoScalarType.FLOATNUMBER: "Number",
            RegoScalarType.BOOL: "Boolean",
            RegoScalarType.DURATION: "Number",
            RegoScalarType.TIME: "Number",
        }
        return f"types.{mapping[self.scalar_type]}{{}}"

    def to_go_rego_types_api_code(self):
        if not self.is_composite:
            return self._to_go_rego_types_api_code()
        return f"types.NewArray([]types.Type{{}}, {self._to_go_rego_types_api_code()})"

class CmdClass(RegoTypeConvertible):
    def __init__(self, go_type, scalar_type, composite):
        super().__init__(go_type, scalar_type, composite)

    @classmethod
    def from_redis_cmd_class(cls, cmd_class):
        mapping = {
            "Cmd": RegoScalarType.ANY,
            "StringCmd": RegoScalarType.STRING,
            "IntCmd": RegoScalarType.INTNUMBER,
            "BoolCmd": RegoScalarType.BOOL,
            "DurationCmd": RegoScalarType.DURATION,
            "StatusCmd": RegoScalarType.STRING,
            "FloatCmd": RegoScalarType.FLOATNUMBER,
            "TimeCmd": RegoScalarType.TIME,
        }
        cmd_class = cmd_class.replace("*", "").strip()
        is_composite = "Slice" in cmd_class
        skip = cmd_class in skip_cmd_classes
        cmd_class = cmd_class.replace("Slice", "").strip()
        scalar_type = mapping.get(cmd_class, None)
        return cls(cmd_class, scalar_type, is_composite), skip

    def to_go_ast_term_return_statement(self, var_name, cmd_name):
        if not self.is_composite:
            return f"""
                {self._to_go_ast_term("term", var_name, cmd_name)} 
                return term, nil"""
        return f"""
        var ret []*ast.Term
        for _, v := range {var_name} {{
            {self._to_go_ast_term("term", "v", cmd_name)}
            ret = append(ret, term)
        }}
        return ast.ArrayTerm(ret...), nil
                """

    def _to_go_ast_term(self, ret_var_name, var_name, cmd_name):
        def ast_ret(term):
            return f"{{ret_var_name}} := ast.{term!s}"

        duration_mapping = {
            "PTTL": ast_ret("IntNumberTerm(int({var_name}.Milliseconds()))"),
            "TTL": ast_ret("IntNumberTerm(int({var_name}.Seconds()))"),
            "OBJECTIDLETIME": ast_ret("IntNumberTerm(int({var_name}.Seconds()))")
        }
        mapping = {
            RegoScalarType.ANY: """
            {ret_var_name} := ast.NullTerm()
            if s, ok := {var_name}.(string); ok {{
                {ret_var_name} = ast.StringTerm(s)
            }}
            """,
            RegoScalarType.STRING: ast_ret("StringTerm({var_name})"),
            RegoScalarType.BOOL: ast_ret("BooleanTerm({var_name})"),
            RegoScalarType.INTNUMBER: ast_ret("IntNumberTerm(int({var_name}))"),
            RegoScalarType.UINTNUMBER: ast_ret("UintNumberTerm(uint({var_name}))"),
            RegoScalarType.FLOATNUMBER: ast_ret("FloatNumberTerm(float64({var_name}))"),
            RegoScalarType.DURATION: duration_mapping,
            RegoScalarType.TIME: ast_ret("IntNumberTerm(int({var_name}.UnixMicro()))")
        }
        if isinstance(mapping[self.scalar_type], dict):
            return f"{mapping[self.scalar_type][cmd_name].format(ret_var_name=ret_var_name, var_name=var_name)}"
        return f"{mapping[self.scalar_type].format(ret_var_name=ret_var_name, var_name=var_name)}"

class ParameterClass(RegoTypeConvertible):
    def __init__(self, go_type, scalar_type, composite):
        super().__init__(go_type, scalar_type, composite)

    @classmethod
    def from_go_type(cls, go_type):
        mapping = {
            "bool": RegoScalarType.BOOL,
            "string": RegoScalarType.STRING,
            "int": RegoScalarType.INTNUMBER,
            "int64": RegoScalarType.INTNUMBER,
            "uint64": RegoScalarType.UINTNUMBER,
            "float64": RegoScalarType.FLOATNUMBER,
            "interface{}": RegoScalarType.ANY,
            "time.Duration": RegoScalarType.DURATION,
            "time.Time": RegoScalarType.TIME,
        }

        is_composite = go_type.startswith("...") or go_type.startswith("[]")
        go_type = go_type.replace("...", "").replace("[]", "").strip()
        scalar_type = mapping[go_type]
        return cls(go_type, scalar_type, is_composite)
    
    def to_go_parse_var_code(self, var_name, idx):
        return  f"""
            {self._to_go_var_declaration(var_name)}
            if err := ast.As(terms[{idx!s}].Value, &{var_name}); err != nil {{
                return nil, err
            }}
            """
    def _to_go_var_declaration(self, var_name):
        if not self.is_composite:
            return f"var {var_name} {self.go_type}"
        return f"var {var_name} []{self.go_type}"

    def to_go_parameter_code(self, var_name, is_last):
        if not self.is_composite or not is_last:
            return self._to_go_parameter_code(var_name)
        return f"{self._to_go_parameter_code(var_name)}..."

    def _to_go_parameter_code(self, var_name):
        if self.scalar_type == RegoScalarType.ANY:
            return f"conva({var_name})" if self.is_composite else f"conv({var_name})"
        return var_name



class CmdSignature:
    @classmethod
    def from_string(cls, signature):
        # print(f"Parsing Signature: {signature!s}", file=sys.stderr)
        cmd_name = re.match(r"(^.*?)\(", signature).group(1)
        signature = signature[len(cmd_name)+1:].strip()
        signature = re.sub(r"ctx context\.Context(, )?", "", signature)
        parameter_types, skip_go_type = CmdSignature.parse_parameter_list(signature[:signature.index(")")].split(","))
        signature = signature[signature.index(")")+1:].strip()
        return_type, skip_cmd_class = CmdClass.from_redis_cmd_class(signature)
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
                parameter_types.insert(0, ParameterClass.from_go_type(go_type))
            except KeyError as e:
                # print(f"KEY ERROR: {e!s}, {p}", file=sys.stderr)
                continue
        return parameter_types, skip
            

    def __init__(self, cmd_name: str, parameter_types: List[ParameterClass], return_type: CmdClass, skip: bool):
        self.cmd_name = cmd_name
        self.parameter_types = parameter_types
        self.return_type = return_type
        self.skip = skip


def main():
    if len(sys.argv) != 2 or sys.argv[1] not in ("doc", "impl"):
        return usage()
    gen_impl = sys.argv[1] == "impl"
    

    if not commands_file.exists():
        print(f"File: {commands_file.absolute()} not found", file=sys.stderr)
        return 1
    if (cmdable := re.search(cmdable_class_regex, commands_file.read_text())) is None:
        print("Could not find Cmdable interface", file=sys.stderr)
        return 1
    command_signatures = cmdable.group(1).split("\n")
    
    if gen_impl:
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
    else:
        print("# Supported Commands:\n")


    register_functions = list()
    for signature in filter(lambda s:s.cmd_name not in skip_command_names, parse_command_signatures(command_signatures)):
        if signature.skip: continue
        register_functions.append(f"register{signature.cmd_name.upper()}")
        # print(f" Generating Function Code for: {signature.cmd_name}")
        if gen_impl:
            code = generate_function_code_for_signature(signature)
            print(code)
        else:
            args = ", ".join(map(lambda p:p.to_go_rego_types_api_code(), signature.parameter_types))
            print(f" - {signature.cmd_name.upper()}( {args!s} ) -> {signature.return_type.to_go_rego_types_api_code()}")

    register_main_func = """
func (p *redisPlugin) registerAutogenCommands() {"""
    for f in register_functions:
        register_main_func += f"""
    {f}(p)"""
    register_main_func += """
}"""
    if gen_impl:
        print(register_main_func)
    else:
        print("\nMore Information about all commands can be found [here](https://redis.io/commands)")

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
        parameter_args.append(p.to_go_parameter_code(var_name, idx==len(signature.parameter_types)-1))
        code += p.to_go_parse_var_code(var_name, idx)
    parameter_args = ",".join(parameter_args)

    code += f"""

            val, err := rdb.{signature.cmd_name}({parameter_args}).Result()
            switch err {{
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                {signature.return_type.to_go_ast_term_return_statement("val", signature.cmd_name.upper())}
            default:
                return nil, err
            }}
        }},
    )
}}
"""
    return code


def usage():
    print(f"Usage: {sys.argv[0]} (impl|doc)")
    return 1


if __name__ == "__main__":
    sys.exit(main())
#!/usr/bin/env python3

import random
import sys
import re
from pathlib import Path
from typing import List, Union

commands_file = Path(__file__).parent.parent / "vendor" / "github.com" / "go-redis" / "redis" / "v8" / "commands.go"
cmdable_class_regex = r"type Cmdable interface {((.|\s)*?)\s}"
skip_command_names = (
    "Pipeline",
    "Pipelined",
    "TxPipelined",
    "TxPipeline",
    "Command",
)
skip_go_types = (
)
skip_cmd_classes = (
)

# TODO: Implement Transactions
# pipe = redis.new_pipe()
# redis.pipe_add(pipe, "incr", ["key"])
# redis.pipe_add(pipe, "expire", ["key", 10])
# [incr_result, expire_result] = redis.pipe_exec(pipe)
#
# [incr_result, expire_result] = redis.piped([
#   {
#       "command": "incr",
#       "args": ["key"]
#   },
#   {
#       "command": "expire",
#       "args": ["key", 10]
#   } 
# ])


class RegoTypeFactory:
    @staticmethod
    def from_go_type(go_type):
        is_pointer = go_type.startswith("*")
        if is_pointer:
            go_type = go_type[1:]

        is_container = go_type.startswith("...") or go_type.startswith("[]")
        go_type = go_type.replace("...", "").replace("[]", "").strip()
        if is_container:
            return RegoArrayType(RegoTypeFactory.from_go_type(go_type), is_pointer=is_pointer)
        scalar_mapping = {
            "bool": RegoBoolType(),
            "string": RegoStringType(),
            "int": RegoIntNumberType(),
            "int64": RegoInt64NumberType(),
            "uint64": RegoUIntNumberType(),
            "float64": RegoFloatNumberType(),
            "interface{}": RegoAnyType(),
            "time.Duration": RegoDurationType(),
            "time.Time": RegoTimeType(),
        }
        redis_mapping = {
            "BitCount": RegoObjType(BitCountObjSpec),
            "LPosArgs": RegoObjType(LPosArgsObjSpec),
            "SetArgs": RegoObjType(SetArgsObjSpec),
            "Sort": RegoObjType(SortObjSpec),
            "GeoLocation": RegoObjType(GeoLocationObjSpec),
            "GeoPos": RegoObjType(GeoPosObjSpec),
            "GeoSearchLocationQuery": RegoObjType(GeoSearchLocationQueryObjSpec),
            "GeoRadiusQuery": RegoObjType(GeoRadiusQueryObjSpec),
            "GeoSearchStoreQuery": RegoObjType(GeoSearchStoreQueryObjSpec),
            "GeoSearchQuery": RegoObjType(GeoSearchQueryObjSpec),
            "Z": RegoObjType(ZObjSpec),
            "ZStore": RegoObjType(ZStoreObjSpec),
            "ZAddArgs": RegoObjType(ZAddArgsObjSpec),
            "ZRangeBy": RegoObjType(ZRangeByObjSpec),
            "ZRangeArgs": RegoObjType(ZRangeArgsObjSpec),
            "XAutoClaimArgs": RegoObjType(XAutoClaimArgsObjSpec),
            "XClaimArgs": RegoObjType(XClaimArgsObjSpec),
            "XAddArgs": RegoObjType(XAddArgsObjSpec),
            "XReadArgs": RegoObjType(XReadArgsObjSpec),
            "XReadGroupArgs": RegoObjType(XReadGroupArgsObjSpec),
            "XPendingExtArgs": RegoObjType(XPendingExtArgsObjSpec),
        }
        if go_type in scalar_mapping:
            return scalar_mapping[go_type].with_go_type(go_type).with_is_pointer(is_pointer)
        return redis_mapping[go_type].with_go_type(f"redis.{go_type}").with_is_pointer(is_pointer)

class RegoType:
    TYPES_API_TYPE = ""
    GO_TYPE = ""

    def __init__(self, go_type=None, is_pointer=False) -> None:
        self._go_type = go_type if go_type is not None else self.GO_TYPE
        self.is_pointer = is_pointer
    
    def with_go_type(self, go_type):
        self._go_type = go_type
        return self
    
    def with_is_pointer(self, is_pointer):
        self.is_pointer = is_pointer
        return self
    
    @property
    def go_type(self):
        if self.is_pointer:
            return f"*{self._go_type}"
        return self._go_type
    
    def to_go_rego_types_api_code(self):
        raise NotImplementedError()

    def to_go_parse_var_code(self, var_name, idx):
        return  f"""
            {self._to_go_var_declaration(var_name)}
            if err := ast.As(terms[{idx!s}].Value, &{var_name}); err != nil {{
                return nil, err
            }}
            """

    def _to_go_var_declaration(self, var_name):
        raise NotImplementedError()
    
    def to_go_parameter_code(self, var_name, is_last):
        raise NotImplementedError()

    def to_go_ast_term_return_statement(self, var_name, cmd_name):
        raise NotImplementedError()
    
    def _to_go_ast_term(self, ret_var_name, var_name, cmd_name):
        raise NotImplementedError()

class RegoScalarType(RegoType):
    def to_go_rego_types_api_code(self):
        return f"types.{self.TYPES_API_TYPE}{{}}"

    def _to_go_var_declaration(self, var_name):
        return f"var {var_name} {self.go_type}"
    

    def to_go_parameter_code(self, var_name, is_last):
        return self._to_go_parameter_code(var_name)

    def _to_go_parameter_code(self, var_name):
        return var_name

    def to_go_ast_term_return_statement(self, var_name, cmd_name):
        return f"""
            {self._to_go_ast_term("term", var_name, cmd_name)} 
            return term, nil"""

    def _to_go_ast_term(self, var_name, cmd_name):
        raise NotImplementedError()
    
class RegoAnyType(RegoScalarType):
    TYPES_API_TYPE = "Any"
    GO_TYPE = "interface{}"

    def _to_go_parameter_code(self, var_name):
        return f"utils.Conv({var_name})"
    
    def _to_go_ast_term(self, ret_var_name, var_name, cmd_name):
            return f"""
            {ret_var_name} := ast.NullTerm()
            if s, ok := {var_name}.(string); ok {{
                {ret_var_name} = ast.StringTerm(s)
            }}
            """

class RegoStringType(RegoScalarType):
    TYPES_API_TYPE = "String"
    GO_TYPE = "string"

    def _to_go_ast_term(self, ret_var_name, var_name, cmd_name):
        return f"{ret_var_name} := ast.StringTerm({var_name})"

class RegoIntNumberType(RegoScalarType):
    TYPES_API_TYPE = "Number"
    GO_TYPE = "int"
    
    def _to_go_ast_term(self, ret_var_name, var_name, cmd_name):
        return f"{ret_var_name} := ast.IntNumberTerm(int({var_name}))"

class RegoInt64NumberType(RegoScalarType):
    TYPES_API_TYPE = "Number"
    GO_TYPE = "int64"
    
    def _to_go_ast_term(self, ret_var_name, var_name, cmd_name):
        return f"{ret_var_name} := ast.IntNumberTerm(int({var_name}))"

class RegoUIntNumberType(RegoScalarType):
    TYPES_API_TYPE = "Number"
    GO_TYPE = "uint64"
    
    def _to_go_ast_term(self, ret_var_name, var_name, cmd_name):
        return f"{ret_var_name} := ast.UIntNumberTerm(uint64({var_name}))"

class RegoFloatNumberType(RegoScalarType):
    TYPES_API_TYPE = "Number"
    GO_TYPE = "float64"
    
    def _to_go_ast_term(self, ret_var_name, var_name, cmd_name):
        return f"{ret_var_name} := ast.FloatNumberTerm(float64({var_name}))"

class RegoBoolType(RegoScalarType):
    TYPES_API_TYPE = "Boolean"
    GO_TYPE = "bool"
    
    def _to_go_ast_term(self, ret_var_name, var_name, cmd_name):
        return f"{ret_var_name} := ast.BooleanTerm({var_name})"

class RegoDurationType(RegoScalarType):
    TYPES_API_TYPE = "Number"
    GO_TYPE = "time.Duration"

    def _to_go_ast_term(self, ret_var_name, var_name, cmd_name):
        duration_mapping = {
            "PTTL": f"{ret_var_name} := ast.IntNumberTerm(int({var_name}.Milliseconds()))",
            "TTL": f"{ret_var_name} := ast.IntNumberTerm(int({var_name}.Seconds()))",
            "OBJECTIDLETIME": f"{ret_var_name} := ast.IntNumberTerm(int({var_name}.Seconds()))",
            "XPENDINGEXT": f"{ret_var_name} := ast.IntNumberTerm(int({var_name}.Milliseconds()))"
        }
        return duration_mapping[cmd_name]


class RegoTimeType(RegoScalarType):
    TYPES_API_TYPE = "Number"
    GO_TYPE = "time.Time"

    def _to_go_ast_term(self, ret_var_name, var_name, cmd_name):
        return f"{ret_var_name} := ast.IntNumberTerm(int({var_name}.UnixMicro()))"


class RegoArrayType(RegoType):
    def __init__(self, containing_type: RegoType, is_pointer=False) -> None:
        self.containing_type = containing_type
        super().__init__(is_pointer=is_pointer)

    def _to_go_var_declaration(self, var_name):
        return f"var {var_name} []{self.containing_type.go_type}"

    def to_go_rego_types_api_code(self):
        return f"types.NewArray([]types.Type{{}}, {self.containing_type.to_go_rego_types_api_code()})"

    def to_go_parameter_code(self, var_name, is_last):
        if not is_last:
            return self._to_go_parameter_code(var_name)
        return f"{self._to_go_parameter_code(var_name)}..."
    
    def _to_go_parameter_code(self, var_name):
        if self.containing_type.TYPES_API_TYPE == "Any":
            return f"utils.Conva({var_name})"
        return var_name

    def to_go_ast_term_return_statement(self, var_name, cmd_name):
        # TODO: handle containing_type == RegoArrayType
        return f"""
        {self._to_go_ast_term("term", var_name, cmd_name)}
        return term,  nil
        """
    
    def _to_go_ast_term(self, ret_var_name, var_name, cmd_name):
        tmp = f"tmp{''.join(random.choices('1234567890', k=5))}"
        code = f"""
        var {tmp} []*ast.Term
        for _, v := range {var_name} {{"""
        if self.containing_type.is_pointer:
            code += f"""
            if v == nil {{
                {tmp} = append({tmp}, ast.NullTerm())
                continue
            }}
            """
        
        code += f"""
            {self.containing_type._to_go_ast_term("term", "v", cmd_name)}
            {tmp} = append({tmp}, term)
        }}
        {ret_var_name} := ast.ArrayTerm({tmp}...)
        """
        return code


class RegoObjType(RegoType):
    def __init__(self, object_spec, is_pointer=False) -> None:
        self.object_spec = object_spec
        super().__init__(is_pointer=is_pointer)
    
    def to_go_rego_types_api_code(self):
        static_properties_list = list()
        for key, value_type in self.object_spec.items():
            static_properties_list.append(f"types.NewStaticProperty(\"{key!s}\", {value_type.to_go_rego_types_api_code()})")
        static_properties_list = ",".join(static_properties_list) 
        return f"types.NewObject([]*types.StaticProperty{{{static_properties_list}}}, nil)"

    def _to_go_var_declaration(self, var_name):
        return f"var {var_name} {self.go_type}"

    def to_go_parameter_code(self, var_name, is_last):
        return var_name

    def to_go_ast_term_return_statement(self, var_name, cmd_name):
        return f"""
        {self._to_go_ast_term("term", var_name, cmd_name)}
        return term,  nil
        """
    
    def _to_go_ast_term(self, ret_var_name, var_name, cmd_name):
        code = ""
        object_args = list()
        for key, value_type in self.object_spec.items():
            object_args.append(key.lower())
            code += f"""
            {value_type._to_go_ast_term(ret_var_name+key, f"{var_name}.{key}", cmd_name)}
            """
        object_args = ",".join(object_args)

        key_value_pairs = list()
        for key, value_type in self.object_spec.items():
            key_value_pairs.append(f"[2]*ast.Term{{ast.StringTerm(\"{key!s}\"), {ret_var_name+key}}}")
        key_value_pairs = ",".join(key_value_pairs)
        code += f"""
        {ret_var_name} := ast.ObjectTerm({key_value_pairs})
        """
        return code
    
class RegoMapType(RegoType):
    def __init__(self, key_type: RegoType, value_type: RegoType, key_var_name="key", value_var_name="value", is_pointer=False):
        self.key_type = key_type
        self.value_type = value_type
        self.key_var_name = key_var_name
        self.value_var_name = value_var_name
        super().__init__(is_pointer=is_pointer)
    
    def to_go_rego_types_api_code(self):
        return f"types.NewObject([]*types.StaticProperty{{}}, types.NewDynamicProperty({self.key_type.to_go_rego_types_api_code()}, {self.value_type.to_go_rego_types_api_code()}))"
    
    def _to_go_var_declaration(self, var_name):
        return f"var {var_name} {self.go_type}"

    def to_go_parameter_code(self, var_name, is_last):
        return var_name

    def to_go_ast_term_return_statement(self, var_name, cmd_name):
        return f"""
        {self._to_go_ast_term("term", var_name, cmd_name)}
        return term,  nil
        """
    
    def _to_go_ast_term(self, ret_var_name, var_name, cmd_name):
        tmp = f"tmp{''.join(random.choices('1234567890', k=5))}"
        code = f"""
        var {tmp} [][2]*ast.Term
        for {self.key_var_name}, {self.value_var_name} := range {var_name} {{
            {self.key_type._to_go_ast_term("k", self.key_var_name, cmd_name)}"""
        if self.value_type.is_pointer:
            code += f"""
            if value == nil {{
                {tmp} = append({tmp}, [2]*ast.Term{{k, ast.NullTerm()}})
                continue
            }}
            """
        
        code += f"""
            {self.value_type._to_go_ast_term("v", self.value_var_name, cmd_name)}
            {tmp} = append({tmp}, [2]*ast.Term{{k, v}})
        }}
        {ret_var_name} := ast.ObjectTerm({tmp}...)
        """
        return code
    

class RegoReturnType: # dont know if this should be RegoType base
    @classmethod
    def from_redis_cmd_type(cls, cmd_type):
        mapping = {
            "StringCmd": StringCmdReturnType,
            "StatusCmd": StatusCmdReturnType,
            "IntCmd": IntCmdReturnType,
            "BoolCmd": BoolCmdReturnType,
            "StringSliceCmd": StringSliceCmdReturnType,
            "DurationCmd": DurationCmdReturnType,
            "FloatCmd": FloatCmdReturnType,
            "SliceCmd": SliceCmdReturnType,
            "IntSliceCmd": IntSliceCmdReturnType,
            "BoolSliceCmd": BoolSliceCmdReturnType,
            "FloatSliceCmd": FloatSliceCmdReturnType,
            "TimeCmd": TimeCmdReturnType,
            "Cmd": CmdReturnType,
            "StringStringMapCmd": StringStringMapCmdReturnType,
            "StringIntMapCmd": StringIntMapCmdReturnType,
            "StringStructMapCmd": StringStructMapCmdReturnType,
            "ScanCmd": ScanCmdReturnType,
            "GeoPosCmd": GeoPosCmdReturnType,
            "GeoSearchLocationCmd": GeoSearchLocationCmdReturnType,
            "GeoLocationCmd": GeoLocationCmdReturnType,
            "ClusterSlotsCmd": ClusterSlotsCmdReturnType,
            "ZSliceCmd": ZSliceCmdReturnType,
            "ZWithKeyCmd": ZWithKeyCmdReturnType,
            "XInfoStreamCmd": XInfoStreamCmdReturnType,
            "XInfoStreamFullCmd": XInfoStreamFullCmdReturnType,
            "XInfoConsumersCmd": XInfoConsumersCmdReturnType,
            "XInfoGroupsCmd": XInfoGroupsCmdReturnType,
            "XAutoClaimJustIDCmd": XAutoClaimJustIDCmdReturnType,
            "XAutoClaimCmd": XAutoClaimCmdReturnType, 
            "XPendingExtCmd": XPendingExtCmdReturnType,
            "XPendingCmd": XPendingCmdReturnType,
            "XStreamSliceCmd": XStreamSliceCmdReturnType,
            "XMessageSliceCmd": XMessageSliceCmdReturnType,
        }
        return mapping[cmd_type.replace("*", "")]

    def __init__(self, return_types: Union[List[RegoType], RegoType]) -> None:
        if isinstance(return_types, RegoType):
            return_types = (return_types,)
        if len(return_types) == 0: raise ValueError("len(return_types) must be > 0.")
        self.return_types = return_types

    @property
    def return_types_count(self):
        return len(self.return_types)

    def to_go_rego_types_api_code(self):
        if self.return_types_count == 1:
            return self.return_types[0].to_go_rego_types_api_code()
        static_types = ",".join([return_type.to_go_rego_types_api_code() for return_type in self.return_types])
        return f"types.NewArray([]types.Type{{{static_types}}}, types.Null{{}})"
    
    def to_go_ast_term_return_statement(self, var_names, cmd_name):
        if self.return_types_count == 1:
            return self.return_types[0].to_go_ast_term_return_statement(var_names[0], cmd_name)
        code = ""
        array_args = list()
        for var_name, return_type in zip(var_names, self.return_types):
            array_args.append(f"t{var_name}")
            code += return_type._to_go_ast_term(f"t{var_name}", var_name, cmd_name)
        array_args = ",".join(array_args)
        code += f"""
        return ast.ArrayTerm({array_args}), nil
        """
        return code

EmptyObjSpec = {}
BitCountObjSpec = {
    "Start": RegoInt64NumberType(),
    "End": RegoInt64NumberType(),
}
LPosArgsObjSpec = {
    "Rank": RegoInt64NumberType(),
    "MaxLen": RegoInt64NumberType(),
}
SetArgsObjSpec = {
    "Mode": RegoStringType(),
    "TTL": RegoDurationType(),
    "ExpireAt": RegoTimeType(),
    "Get": RegoBoolType(),
    "KeepTTL": RegoBoolType(),
}
SortObjSpec = {
    "By": RegoStringType(),
    "Offset": RegoInt64NumberType(),
    "Count": RegoInt64NumberType(),
    "Get": RegoArrayType(RegoStringType()),
    "Order": RegoStringType(),
    "Alpha": RegoBoolType()
}
GeoLocationObjSpec = {
    "Name": RegoStringType(),
    "Longitude": RegoFloatNumberType(),
    "Latitude": RegoFloatNumberType(),
    "Dist": RegoFloatNumberType(),
    "GeoHash": RegoIntNumberType(),
}
GeoPosObjSpec = {
    "Longitude": RegoFloatNumberType(),
    "Latitude": RegoFloatNumberType(),
}
GeoSearchLocationQueryObjSpec = {
    "WithCoord": RegoBoolType(),
    "WithDist": RegoBoolType(),
    "WithHash": RegoBoolType(),
}
GeoRadiusQueryObjSpec = {
    "Radius": RegoFloatNumberType(),
    "Unit": RegoStringType(),
    "WithCoord": RegoBoolType(),
    "WithDist": RegoBoolType(),
    "WithGeoHash": RegoBoolType(),
    "Count": RegoIntNumberType(),
    "Sort": RegoStringType(),
    "Store": RegoStringType(),
    "StoreDist": RegoStringType()
}
GeoSearchStoreQueryObjSpec = {
    "StoreDist": RegoBoolType()
}
GeoSearchQueryObjSpec = {
    "Member": RegoStringType(),
    "Longitude": RegoFloatNumberType(),
    "Latitude": RegoFloatNumberType(),
    "Radius": RegoFloatNumberType(),
    "RadiusUnit": RegoStringType(),
    "BoxWidth": RegoFloatNumberType(),
    "BoxHeight": RegoFloatNumberType(),
    "BoxUnit": RegoStringType(),
    "Sort": RegoStringType(),
    "Count": RegoIntNumberType(),
    "CountAny": RegoBoolType(),
}
ClusterNodeObjSpec = {
    "ID": RegoStringType(),
    "Addr": RegoStringType()
}
ClusterSlotObjSpec = {
    "Start": RegoIntNumberType(),
    "End": RegoIntNumberType(),
    "Nodes": RegoArrayType(RegoObjType(ClusterNodeObjSpec))
}
ZObjSpec = {
    "Score": RegoFloatNumberType(),
    "Member": RegoAnyType(),
}
ZStoreObjSpec = {
    "Keys": RegoArrayType(RegoStringType()),
    "Weights": RegoArrayType(RegoFloatNumberType()), 
    "Aggregate": RegoStringType()
}
ZAddArgsObjSpec = {
    "NX": RegoBoolType(),
    "XX": RegoBoolType(),
    "LT": RegoBoolType(),
    "GT": RegoBoolType(),
    "Ch": RegoBoolType(),
    "Members": RegoArrayType(RegoObjType(ZObjSpec)),
}
ZWithKeyObjSpec = {
    "Key": RegoStringType()
}
ZRangeByObjSpec = {
    "Min": RegoStringType(),
    "Max": RegoStringType(),
    "Offset": RegoInt64NumberType(),
    "Count": RegoInt64NumberType(),
}
ZRangeArgsObjSpec = {
    "Key": RegoStringType(),
    "Start": RegoAnyType(),
    "Stop": RegoAnyType(),
    "ByScore": RegoBoolType(),
    "ByLex": RegoBoolType(),
    "Rev": RegoBoolType(),
    "Offset": RegoInt64NumberType(),
    "Count": RegoInt64NumberType(),
}
XInfoConsumerObjSpec = {
    "Name": RegoStringType(),
    "Pending": RegoInt64NumberType(),
    "Idle": RegoInt64NumberType()
}
XMessageObjSpec = {
    "ID": RegoStringType(),
    "Values": RegoMapType(RegoStringType(), RegoAnyType())
}
XInfoStreamObjSpec = {
    "Length": RegoInt64NumberType(),
    "RadixTreeKeys": RegoInt64NumberType(),
    "RadixTreeNodes": RegoInt64NumberType(),
    "Groups": RegoInt64NumberType(),
    "LastGeneratedID": RegoStringType(),
    "FirstEntry": RegoObjType(XMessageObjSpec),
    "LastEntry": RegoObjType(XMessageObjSpec),
}
XInfoStreamGroupPendingObjSpec = {
    "ID": RegoStringType(),
    "Consumer": RegoStringType(),
    "DeliveryTime": RegoTimeType(),
    "DeliveryCount": RegoInt64NumberType(),
}
XInfoStreamConsumerPendingObjSpec = {
    "ID": RegoStringType(),
    "DeliveryTime": RegoTimeType(),
    "DeliveryCount": RegoInt64NumberType(),
}
XInfoStreamConsumerObjSpec = {
    "Name": RegoStringType(),
    "SeenTime": RegoTimeType(),
    "PelCount": RegoInt64NumberType(),
    "Pending": RegoArrayType(RegoObjType(XInfoStreamConsumerPendingObjSpec)),
}
XInfoStreamGroupObjSpec = {
    "Name": RegoStringType(),
    "LastDeliveredID": RegoStringType(),
    "PelCount": RegoInt64NumberType(),
    "Pending": RegoArrayType(RegoObjType(XInfoStreamGroupPendingObjSpec)),
    "Consumers": RegoArrayType(RegoObjType(XInfoStreamConsumerObjSpec))
}
XInfoStreamFullObjSpec = {
    "Length": RegoInt64NumberType(),
    "RadixTreeKeys": RegoInt64NumberType(),
    "RadixTreeNodes": RegoInt64NumberType(),
    "LastGeneratedID": RegoStringType(),
    "Entries": RegoArrayType(RegoObjType(XMessageObjSpec)),
    "Groups": RegoArrayType(RegoObjType(XInfoStreamGroupObjSpec))
}
XInfoGroupObjSpec = {
    "Name": RegoStringType(),
    "Consumers": RegoInt64NumberType(),
    "Pending": RegoInt64NumberType(),
    "LastDeliveredID": RegoStringType(),
}
XAutoClaimArgsObjSpec = {
    "Stream": RegoStringType(),
    "Group": RegoStringType(),
    "MinIdle": RegoDurationType(),
    "Start": RegoStringType(),
    "Count": RegoInt64NumberType(),
    "Consumer": RegoStringType(),
}
XClaimArgsObjSpec = {
    "Stream": RegoStringType(),
    "Group": RegoStringType(),
    "Consumer": RegoStringType(),
    "MinIdle": RegoDurationType(),
    "Messages": RegoArrayType(RegoStringType())
}
XAddArgsObjSpec = {
     "Stream": RegoStringType(),
     "NoMkStream": RegoBoolType(),
     "MaxLen": RegoInt64NumberType(),
     "MaxLenApprox": RegoInt64NumberType(),
     "MinID": RegoStringType(),
     "Approx": RegoBoolType(),
     "Limit": RegoInt64NumberType(),
     "ID": RegoStringType(),
     "Values": RegoAnyType(),
}
XPendingExtObjSpec = {
    "ID": RegoStringType(),
    "Consumer": RegoStringType(),
    "Idle": RegoDurationType(),
    "RetryCount": RegoInt64NumberType(),
}
XPendingExtArgsObjSpec = {
    "Stream": RegoStringType(),
    "Group": RegoStringType(),
    "Idle": RegoDurationType(),
    "Start": RegoStringType(),
    "End": RegoStringType(),
    "Count": RegoInt64NumberType(),
    "Consumer": RegoStringType(),
}
XPendingObjSpec = {
    "Count": RegoInt64NumberType(),
    "Lower": RegoStringType(),
    "Higher": RegoStringType(),
    "Consumers": RegoMapType(RegoStringType(), RegoInt64NumberType()),
}
XStreamObjSpec = {
    "Stream": RegoStringType(),
    "Messages": RegoArrayType(RegoObjType(XMessageObjSpec)),
}
XReadArgsObjSpec = {
    "Streams": RegoArrayType(RegoStringType()),
    "Count": RegoInt64NumberType(),
    "Block": RegoDurationType(),
}
XReadGroupArgsObjSpec = {
    "Group": RegoStringType(),
    "Consumer": RegoStringType(),
    "Streams": RegoArrayType(RegoStringType()),
    "Count": RegoInt64NumberType(),
    "Block": RegoDurationType(),
    "NoAck": RegoBoolType(),
}

StringCmdReturnType = RegoReturnType(RegoStringType())
StatusCmdReturnType = RegoReturnType(RegoStringType())
IntCmdReturnType = RegoReturnType(RegoIntNumberType())
BoolCmdReturnType = RegoReturnType(RegoBoolType())
StringSliceCmdReturnType = RegoReturnType(RegoArrayType(RegoStringType()))
DurationCmdReturnType = RegoReturnType(RegoDurationType())
FloatCmdReturnType = RegoReturnType(RegoFloatNumberType())
SliceCmdReturnType = RegoReturnType(RegoArrayType(RegoAnyType()))
IntSliceCmdReturnType = RegoReturnType(RegoArrayType(RegoIntNumberType()))
BoolSliceCmdReturnType = RegoReturnType(RegoArrayType(RegoBoolType()))
FloatSliceCmdReturnType = RegoReturnType(RegoArrayType(RegoFloatNumberType()))
TimeCmdReturnType = RegoReturnType(RegoTimeType())
StringStringMapCmdReturnType = RegoReturnType(RegoMapType(RegoStringType(), RegoStringType()))
StringIntMapCmdReturnType = RegoReturnType(RegoMapType(RegoStringType(), RegoInt64NumberType()))
StringStructMapCmdReturnType = RegoReturnType(RegoMapType(RegoStringType(), RegoObjType(EmptyObjSpec), value_var_name="_"))
CmdReturnType = RegoReturnType(RegoAnyType())
ScanCmdReturnType = RegoReturnType((RegoArrayType(RegoStringType()), RegoUIntNumberType()))
GeoPosCmdReturnType = RegoReturnType(RegoArrayType(RegoObjType(GeoPosObjSpec, is_pointer=True)))
GeoSearchLocationCmdReturnType = RegoReturnType(RegoArrayType(RegoObjType(GeoLocationObjSpec)))
GeoLocationCmdReturnType = RegoReturnType(RegoArrayType(RegoObjType(GeoLocationObjSpec)))
ClusterSlotsCmdReturnType = RegoReturnType(RegoArrayType(RegoObjType(ClusterSlotObjSpec)))
ZSliceCmdReturnType = RegoReturnType(RegoArrayType(RegoObjType(ZObjSpec)))
ZWithKeyCmdReturnType = RegoReturnType(RegoObjType(ZWithKeyObjSpec, is_pointer=True))
XInfoConsumersCmdReturnType = RegoReturnType(RegoArrayType(RegoObjType(XInfoConsumerObjSpec)))
XInfoStreamCmdReturnType = RegoReturnType(RegoObjType(XInfoStreamObjSpec, is_pointer=True))
XInfoStreamFullCmdReturnType = RegoReturnType(RegoObjType(XInfoStreamFullObjSpec, is_pointer=True))
XInfoGroupsCmdReturnType = RegoReturnType(RegoArrayType(RegoObjType(XInfoGroupObjSpec)))
XAutoClaimJustIDCmdReturnType = RegoReturnType((RegoArrayType(RegoStringType()), RegoStringType()))
XAutoClaimCmdReturnType = RegoReturnType((RegoArrayType(RegoObjType(XMessageObjSpec)), RegoStringType()))
XPendingExtCmdReturnType = RegoReturnType(RegoArrayType(RegoObjType(XPendingExtObjSpec)))
XPendingCmdReturnType = RegoReturnType(RegoObjType(XPendingObjSpec ,is_pointer=True))
XStreamSliceCmdReturnType = RegoReturnType(RegoArrayType(RegoObjType(XStreamObjSpec)))
XMessageSliceCmdReturnType = RegoReturnType(RegoArrayType(RegoObjType(XMessageObjSpec)))


class CmdSignature:
    @classmethod
    def from_string(cls, signature):
        # print(f"Parsing Signature: {signature!s}", file=sys.stderr)
        cmd_name = re.match(r"(^.*?)\(", signature)
        if cmd_name is None or cmd_name.group(1) is None:
            return cls(None, None, None, True) 
        cmd_name = cmd_name.group(1)
        signature = signature[len(cmd_name)+1:].strip()
        signature = re.sub(r"ctx context\.Context(, )?", "", signature)
        parameter_types, skip_go_type = CmdSignature.parse_parameter_list(signature[:signature.index(")")].split(","))
        signature = signature[signature.index(")")+1:].strip()
        if skip_go_type or signature in skip_cmd_classes or cmd_name in skip_command_names:
            return cls(None, None, None, True)
        return_type = RegoReturnType.from_redis_cmd_type(signature)
        return cls(cmd_name, parameter_types, return_type, False)
    
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
                parameter_types.insert(0, RegoTypeFactory.from_go_type(go_type))
            except KeyError as e:
                continue
        return parameter_types, skip
            

    def __init__(self, cmd_name: str, parameter_types: List[RegoType], return_type: RegoReturnType, skip: bool):
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
        print("package redisManager")
        print("""
import (
    "time"
    "github.com/tibotix/opa-redis-plugin/utils"

    "github.com/go-redis/redis/v8"
    "github.com/open-policy-agent/opa/ast"
    "github.com/open-policy-agent/opa/rego"
    "github.com/open-policy-agent/opa/types"
)
""")
    else:
        print("# Supported Commands:\n")
        print("All function signatures are based on the corresponding function of the [go-redis](https://github.com/go-redis/redis) library. Check out their documentation for more details.")
        print("\nMore General Information about all commands can be found on the official [redis site](https://redis.io/commands)\n\n")


    register_functions = list()
    for signature in filter(lambda s:not s.skip, parse_command_signatures(command_signatures)):
        register_functions.append(f"register{signature.cmd_name.upper()}")
        print(f" Generating Function Code for: {signature.cmd_name}", file=sys.stderr)
        if gen_impl:
            code = generate_function_code_for_signature(signature)
            print(code)
        else:
            args = ", ".join(map(lambda p:p.to_go_rego_types_api_code(), signature.parameter_types))
            print(f" - `{signature.cmd_name.upper()}( {args!s} ) -> {signature.return_type.to_go_rego_types_api_code()}`")

    register_main_func = """
func (m *RedisManager) registerAutogenCommands() {"""
    for f in register_functions:
        register_main_func += f"""
    {f}(m)"""
    register_main_func += """
}"""
    if gen_impl:
        print(register_main_func)

    return 0


def parse_command_signatures(command_signatures):
    for signature in map(lambda x:x.strip(), filter(lambda x:x and "//" not in x, command_signatures)):
        yield CmdSignature.from_string(signature)


def generate_function_code_for_signature(signature: CmdSignature):
    args = ",".join(map(lambda p:p.to_go_rego_types_api_code(), signature.parameter_types))

    code = f"""
func register{signature.cmd_name.upper()}(m *RedisManager) {{
    rego.RegisterBuiltinDyn(
        &rego.Function{{
            Name: \"redis.{signature.cmd_name.lower()}\",
            Decl: types.NewFunction(types.Args({args}), {signature.return_type.to_go_rego_types_api_code()}),
            Memoize: true,
            Nondeterministic: true,
        }},
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {{
            rdb, err := m.RedisProxy.Get()
            if err != nil {{
                return nil, err
            }}

"""

    parameter_args = ["m.RedisContext"]
    for idx, p in enumerate(signature.parameter_types):
        var_name = f"v{idx!s}"
        parameter_args.append(p.to_go_parameter_code(var_name, idx==len(signature.parameter_types)-1))
        code += p.to_go_parse_var_code(var_name, idx)
    parameter_args = ",".join(parameter_args)

    var_names = list(map(lambda c:f"r{c!s}", range(signature.return_type.return_types_count)))
    res_args = ",".join(var_names)

    code += f"""

            val := rdb.{signature.cmd_name}({parameter_args})
            {res_args} := val.Val()
            switch err := val.Err(); err {{
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                {signature.return_type.to_go_ast_term_return_statement(var_names, signature.cmd_name.upper())}
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
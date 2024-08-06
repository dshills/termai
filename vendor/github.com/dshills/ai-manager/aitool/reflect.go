package aitool

import (
	"fmt"
	"reflect"
	"strings"
)

func convertAndValidateType(name, desc string, paramType string) (Property, error) {
	switch paramType {
	case "string":
		return NewString(name, desc, true), nil
	case "nil":
		return NewNumeric(name, desc, true, NumericNil), nil
	case "bool":
		return NewNumeric(name, desc, true, NumericBool), nil
	case "int":
		return NewNumeric(name, desc, true, NumericInt), nil
	case "int8":
		return NewNumeric(name, desc, true, NumericInt), nil
	case "int16":
		return NewNumeric(name, desc, true, NumericInt), nil
	case "int32":
		return NewNumeric(name, desc, true, NumericInt), nil
	case "int64":
		return NewNumeric(name, desc, true, NumericInt), nil
	case "uint":
		return NewNumeric(name, desc, true, NumericInt), nil
	case "uint8":
		return NewNumeric(name, desc, true, NumericInt), nil
	case "uint16":
		return NewNumeric(name, desc, true, NumericInt), nil
	case "uint32":
		return NewNumeric(name, desc, true, NumericInt), nil
	case "uint64":
		return NewNumeric(name, desc, true, NumericInt), nil
	case "float32":
		return NewNumeric(name, desc, true, NumericFloat), nil
	case "float64":
		return NewNumeric(name, desc, true, NumericFloat), nil
	}
	return nil, fmt.Errorf("%v invalid type", paramType)
}

// getFunctionSchema takes an arbitrary function and returns its JSON schema.
func getFunctionParamTypes(fn interface{}) ([]string, error) {
	// Validate that fn is a function
	fnType := reflect.TypeOf(fn)
	if fnType.Kind() != reflect.Func {
		return nil, fmt.Errorf("expected a function, got %v", fnType)
	}

	props := []string{}

	// Populate the parameters
	for i := 0; i < fnType.NumIn(); i++ {
		paramType := fnType.In(i)
		props = append(props, paramType.Name())
	}

	return props, nil
}

func splitInfo(info string) (string, string) {
	splits := strings.Split(info, ":")
	if len(splits) < 2 {
		return info, ""
	}
	return strings.TrimSpace(splits[0]), strings.TrimSpace(splits[1])
}

// ToolFromFunc will convert a function into a Tool for simulated calling by an AI
// The paramInfo is in the form "paramName:paramDescription" if a description is blank it is ignored
// Param and function names and descriptions are important to tell the AI what they should use to call
// Currently does not support building arrays or structs
func ToolFromFunc(fnName, fnDesc string, fn interface{}, paramInfo ...string) (*Tool, error) {
	ptypes, err := getFunctionParamTypes(fn)
	if err != nil {
		return nil, err
	}
	props := []Property{}
	for i := range ptypes {
		name := fmt.Sprintf("param%v", i)
		desc := ""
		if len(paramInfo) > i {
			name, desc = splitInfo(paramInfo[i])
		}
		prop, err := convertAndValidateType(name, desc, ptypes[i])
		if err != nil {
			return nil, err
		}
		props = append(props, prop)
	}
	tool := NewTool(fnName, fnDesc, props...)
	return tool, nil
}

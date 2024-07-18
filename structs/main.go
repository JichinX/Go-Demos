package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var defaultsMap = map[reflect.Kind]any{}
var (
	defaultString string
	defaultBool   bool
	defaultInt    int
	defaultUint   uint
	defaultFloat  float64
	defaultSlice  []string
	defaultMap    map[string]any
	defaultStruct any
	defaultArray  []any
)

func init() {
	defaultsMap[reflect.String] = defaultString
	defaultsMap[reflect.Bool] = defaultBool
	defaultsMap[reflect.Int] = defaultInt
	defaultsMap[reflect.Uint] = defaultUint
	defaultsMap[reflect.Float64] = defaultFloat
	defaultsMap[reflect.Slice] = defaultSlice
	defaultsMap[reflect.Map] = defaultMap
	defaultsMap[reflect.Struct] = defaultStruct
}

func StructToMap(obj any, omitempty bool) (result map[string]any, err error) {
	t := reflect.TypeOf(obj)
	switch t.Kind() {
	case reflect.Struct:
		return toMapForStruct(obj, omitempty)
	case reflect.Pointer:
		return toMapForPtr(obj, omitempty)
	default:
		return nil, errors.New(fmt.Sprintf("unsupported type: %v", t.Kind()))
		//case reflect.Invalid:
		//case reflect.Bool:
		//case reflect.Int:
		//case reflect.Int8:
		//case reflect.Int16:
		//case reflect.Int32:
		//case reflect.Int64:
		//case reflect.Uint:
		//case reflect.Uint8:
		//case reflect.Uint16:
		//case reflect.Uint32:
		//case reflect.Uint64:
		//case reflect.Uintptr:
		//case reflect.Float32:
		//case reflect.Float64:
		//case reflect.Complex64:
		//case reflect.Complex128:
		//case reflect.Array:
		//case reflect.Chan:
		//case reflect.Func:
		//case reflect.Interface:
		//case reflect.Map:
		//case reflect.Slice:
		//case reflect.String:
		//case reflect.UnsafePointer:

	}

}

func toMapForPtr(obj any, omitempty bool) (result map[string]any, err error) {
	return nil, errors.New("unsupported type ptr")
}

// toMapForStruct
func toMapForStruct(obj any, omitempty bool) (result map[string]any, err error) {

	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	result = make(map[string]any)

	for i := range t.NumField() {
		if t.Field(i).PkgPath != "" {
			continue
		}
		tag, have := t.Field(i).Tag.Lookup("json")
		var realName string
		if have {
			jsonTag := strings.Split(tag, ",")
			realName = jsonTag[0]
		} else {
			realName = t.Field(i).Name
		}

		vv := v.Field(i).Interface()
		vvKind := v.Field(i).Kind()
		if omitempty && isDefaultVal(vvKind, vv) {
			// 如果忽略默认值。且值为默认值
			continue
		}
		switch vvKind {
		case reflect.Struct:
			structToMap, e := StructToMap(vv, omitempty)
			if e != nil {
				return nil, fmt.Errorf("failed to convert map for value: %v \n %w", vv, e)
			}
			result[realName] = structToMap
		default:
			result[realName] = vv
		}
	}
	return
}

func isDefaultVal(kind reflect.Kind, vv any) bool {
	return defaultsMap[kind] == vv
}

func toJson(obj any) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return ""
	} else {
		return string(bytes)
	}
}

func main() {
	var data = struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
		City string `json:"city,omitempty"`
	}{
		Id:   1,
		Name: "",
		Age:  0,
		City: "shanghai",
	}

	result, err := StructToMap(data, true)

	if err != nil {
		fmt.Println(err)
	}
	var sql string
	var args []any
	for k, v := range result {
		sql += k + " = ? "
		args = append(args, v)
	}
	fmt.Println("sql---> ", sql)
	fmt.Println("args---> ", args)
}

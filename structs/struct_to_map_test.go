package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name   string `json:"name,omitempty"`
	Age    int    `json:"age"`
	Gender string `json:"gender,omitempty"`
	City   string `json:"city"`
}
type Employee struct {
	Name    string   `json:"name,omitempty"`
	Persons []Person `json:"persons,omitempty"`
	Ages    []int    `json:"ages,omitempty"`
}

func Test_toMapForStruct(t *testing.T) {

	pObj := Person{Name: "xujc", Age: 30, Gender: "男", City: "青岛"}
	pWant := map[string]any{"name": "xujc", "age": 30, "gender": "男", "city": "青岛"}
	ppObj := Employee{Name: "xujc", Persons: []Person{pObj, pObj}, Ages: []int{1, 2, 3}}
	ppWant := map[string]any{
		"name":    "xujc",
		"persons": []Person{pObj, pObj},
		"ages":    []int{1, 2, 3},
	}

	tests := []struct {
		name       string
		args       any
		wantResult map[string]any
		wantErr    bool
	}{
		{
			"测试 1",
			pObj,
			pWant,
			false,
		},
		{
			"测试 2",
			ppObj,
			ppWant,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := toMapForStruct(tt.args, false)
			if (err != nil) != tt.wantErr {
				t.Errorf("toMapForStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(toJson(gotResult), toJson(tt.wantResult)) {
				t.Errorf("toMapForStruct() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

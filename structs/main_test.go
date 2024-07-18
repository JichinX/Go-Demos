package main

import (
	"reflect"
	"testing"
)

func Test_toMapForStructOmitempty(t *testing.T) {

	pObj := Person{Gender: "男", City: "青岛"}
	pWant := map[string]any{"gender": "男", "city": "青岛"}

	type args struct {
		obj       any
		omitempty bool
	}
	tests := []struct {
		name       string
		args       args
		wantResult map[string]any
		wantErr    bool
	}{
		{
			"测试",
			args{pObj, true},
			pWant,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := toMapForStruct(tt.args.obj, tt.args.omitempty)
			if (err != nil) != tt.wantErr {
				t.Errorf("toMapForStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("toMapForStruct() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

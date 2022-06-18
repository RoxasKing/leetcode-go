package main

import (
	"reflect"
	"testing"
)

func Test_duplicateZeros(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 0, 2, 3, 0, 4, 5, 0}}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{"2", args{[]int{1, 2, 3}}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if duplicateZeros(tt.args.arr); !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("duplicateZeros() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

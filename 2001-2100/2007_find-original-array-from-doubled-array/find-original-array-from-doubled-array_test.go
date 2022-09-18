package main

import (
	"reflect"
	"testing"
)

func Test_findOriginalArray(t *testing.T) {
	type args struct {
		changed []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 3, 4, 2, 6, 8}}, []int{1, 3, 4}},
		{"2", args{[]int{6, 3, 0, 1}}, nil},
		{"3", args{[]int{1}}, nil},
		{"4", args{[]int{0}}, nil},
		{"5", args{[]int{0, 0}}, []int{0}},
		{"6", args{[]int{0, 0, 0, 0}}, []int{0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOriginalArray(tt.args.changed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOriginalArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

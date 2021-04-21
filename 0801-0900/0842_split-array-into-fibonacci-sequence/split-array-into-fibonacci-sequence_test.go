package main

import (
	"reflect"
	"testing"
)

func Test_splitIntoFibonacci(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{"123456579"}, []int{123, 456, 579}},
		{"2", args{"11235813"}, []int{1, 1, 2, 3, 5, 8, 13}},
		{"3", args{"112358130"}, []int{}},
		{"4", args{"0123"}, []int{}},
		{"5", args{"1101111"}, []int{11, 0, 11, 11}},
		{
			"6",
			args{"539834657215398346785398346991079669377161950407626991734534318677529701785098211336528511"},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitIntoFibonacci(tt.args.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitIntoFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

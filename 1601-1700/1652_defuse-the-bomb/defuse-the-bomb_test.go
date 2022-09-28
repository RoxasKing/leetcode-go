package main

import (
	"reflect"
	"testing"
)

func Test_decrypt(t *testing.T) {
	type args struct {
		code []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{5, 7, 1, 4}, 3}, []int{12, 10, 16, 13}},
		{"2", args{[]int{1, 2, 3, 4}, 0}, []int{0, 0, 0, 0}},
		{"3", args{[]int{2, 4, 9, 3}, -2}, []int{12, 5, 6, 13}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decrypt(tt.args.code, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

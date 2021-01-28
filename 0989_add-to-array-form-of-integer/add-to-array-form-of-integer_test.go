package main

import (
	"reflect"
	"testing"
)

func Test_addToArrayForm(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 2, 0, 0}, 34}, []int{1, 2, 3, 4}},
		{"2", args{[]int{2, 7, 4}, 181}, []int{4, 5, 5}},
		{"3", args{[]int{2, 1, 5}, 806}, []int{1, 0, 2, 1}},
		{"4", args{[]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 1}, []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addToArrayForm(tt.args.A, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addToArrayForm() = %v, want %v", got, tt.want)
			}
		})
	}
}

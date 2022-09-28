package main

import (
	"reflect"
	"testing"
)

func Test_missingTwo(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1}}, []int{2, 3}},
		{"2", args{[]int{2, 3}}, []int{1, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingTwo(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("missingTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

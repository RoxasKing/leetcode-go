package main

import (
	"reflect"
	"testing"
)

func Test_sortArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{5, 2, 3, 1}}, []int{1, 2, 3, 5}},
		{"2", args{[]int{5, 1, 1, 2, 0, 0}}, []int{0, 0, 1, 1, 2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortArray2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{5, 2, 3, 1}}, []int{1, 2, 3, 5}},
		{"2", args{[]int{5, 1, 1, 2, 0, 0}}, []int{0, 0, 1, 1, 2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray2() = %v, want %v", got, tt.want)
			}
		})
	}
}

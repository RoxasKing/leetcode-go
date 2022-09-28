package main

import "testing"

func Test_canFormArray(t *testing.T) {
	type args struct {
		arr    []int
		pieces [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{15, 88}, [][]int{{88}, {15}}}, true},
		{"2", args{[]int{49, 18, 6}, [][]int{{16, 18, 49}}}, false},
		{"3", args{[]int{91, 4, 64, 78}, [][]int{{78}, {4, 64}, {91}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFormArray(tt.args.arr, tt.args.pieces); got != tt.want {
				t.Errorf("canFormArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

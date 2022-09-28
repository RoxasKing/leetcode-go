package main

import "testing"

func Test_findLength(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}}, 3},
		{"2", args{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}}, 5},
		{"3", args{[]int{3, 2, 1, 4, 7}, []int{1, 2, 3, 2, 1}}, 3},
		{"4", args{[]int{1, 1, 0, 0, 1}, []int{1, 1, 0, 0, 0}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLength(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("findLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

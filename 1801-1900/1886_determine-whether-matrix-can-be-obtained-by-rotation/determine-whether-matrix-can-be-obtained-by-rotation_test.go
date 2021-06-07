package main

import "testing"

func Test_findRotation(t *testing.T) {
	type args struct {
		mat    [][]int
		target [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[][]int{{0, 1}, {1, 0}}, [][]int{{1, 0}, {0, 1}}}, true},
		{"2", args{[][]int{{0, 1}, {1, 1}}, [][]int{{1, 0}, {0, 1}}}, false},
		{"3", args{[][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}, [][]int{{1, 1, 1}, {0, 1, 0}, {0, 0, 0}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRotation(tt.args.mat, tt.args.target); got != tt.want {
				t.Errorf("findRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}

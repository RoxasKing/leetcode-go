package main

import "testing"

func Test_domino(t *testing.T) {
	type args struct {
		n      int
		m      int
		broken [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{2, 3, [][]int{{1, 0}, {1, 1}}}, 2},
		{"2", args{3, 3, [][]int{}}, 4},
		{"3", args{3, 3, [][]int{{1, 1}, {1, 2}}}, 3},
		{"4", args{4, 4, [][]int{{1, 0}, {1, 3}}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := domino(tt.args.n, tt.args.m, tt.args.broken); got != tt.want {
				t.Errorf("domino() = %v, want %v", got, tt.want)
			}
		})
	}
}

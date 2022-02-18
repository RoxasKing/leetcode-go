package main

import "testing"

func Test_checkWays(t *testing.T) {
	type args struct {
		pairs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{1, 2}, {2, 3}}}, 1},
		{"2", args{[][]int{{1, 2}, {2, 3}, {1, 3}}}, 2},
		{"3", args{[][]int{{1, 2}, {2, 3}, {2, 4}, {1, 5}}}, 0},
		{"4", args{[][]int{{1, 5}, {1, 3}, {2, 3}, {2, 4}, {3, 5}, {3, 4}}}, 2},
		{"5", args{[][]int{{3, 5}, {4, 5}, {2, 5}, {1, 5}, {1, 4}, {2, 4}}}, 1},
		{"6", args{[][]int{{4, 5}, {3, 4}, {2, 4}}}, 1},
		{"7", args{[][]int{{3, 4}, {2, 3}, {4, 5}, {2, 4}, {2, 5}, {1, 5}, {1, 4}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkWays(tt.args.pairs); got != tt.want {
				t.Errorf("checkWays() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_calculateMinimumHP(t *testing.T) {
	type args struct {
		dungeon [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}}, 7},
		{"2", args{[][]int{{0}}}, 1},
		{"3", args{[][]int{{0, 5}, {-2, -3}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMinimumHP(tt.args.dungeon); got != tt.want {
				t.Errorf("calculateMinimumHP() = %v, want %v", got, tt.want)
			}
		})
	}
}

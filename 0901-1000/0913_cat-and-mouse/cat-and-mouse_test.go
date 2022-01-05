package main

import "testing"

func Test_catMouseGame(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{2, 5}, {3}, {0, 4, 5}, {1, 4, 5}, {2, 3}, {0, 2, 3}}}, 0},
		{"2", args{[][]int{{1, 3}, {0}, {3}, {0, 2}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := catMouseGame(tt.args.graph); got != tt.want {
				t.Errorf("catMouseGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

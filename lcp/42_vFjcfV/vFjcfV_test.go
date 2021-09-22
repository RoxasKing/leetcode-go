package main

import "testing"

func Test_circleGame(t *testing.T) {
	type args struct {
		toys    [][]int
		circles [][]int
		r       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{3, 3, 1}, {3, 2, 1}}, [][]int{{4, 3}}, 2}, 1},
		{"2", args{[][]int{{1, 3, 2}, {4, 3, 1}, {7, 1, 2}}, [][]int{{1, 0}, {3, 3}}, 4}, 2},
		{"3", args{[][]int{{3, 4, 5}, {1, 4, 4}, {4, 4, 1}, {1, 5, 5}}, [][]int{{4, 1}, {4, 2}}, 6}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := circleGame(tt.args.toys, tt.args.circles, tt.args.r); got != tt.want {
				t.Errorf("circleGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

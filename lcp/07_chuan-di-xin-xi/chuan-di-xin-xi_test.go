package main

import (
	"testing"
)

func Test_numWays(t *testing.T) {
	type args struct {
		n        int
		relation [][]int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5, [][]int{{0, 2}, {2, 1}, {3, 4}, {2, 3}, {1, 4}, {2, 0}, {0, 4}}, 3}, 3},
		{"2", args{3, [][]int{{0, 2}, {2, 1}}, 2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWays(tt.args.n, tt.args.relation, tt.args.k); got != tt.want {
				t.Errorf("numWays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numWays2(t *testing.T) {
	type args struct {
		n        int
		relation [][]int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5, [][]int{{0, 2}, {2, 1}, {3, 4}, {2, 3}, {1, 4}, {2, 0}, {0, 4}}, 3}, 3},
		{"2", args{3, [][]int{{0, 2}, {2, 1}}, 2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWays2(tt.args.n, tt.args.relation, tt.args.k); got != tt.want {
				t.Errorf("numWays2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numWays3(t *testing.T) {
	type args struct {
		n        int
		relation [][]int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5, [][]int{{0, 2}, {2, 1}, {3, 4}, {2, 3}, {1, 4}, {2, 0}, {0, 4}}, 3}, 3},
		{"2", args{3, [][]int{{0, 2}, {2, 1}}, 2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWays3(tt.args.n, tt.args.relation, tt.args.k); got != tt.want {
				t.Errorf("numWays3() = %v, want %v", got, tt.want)
			}
		})
	}
}

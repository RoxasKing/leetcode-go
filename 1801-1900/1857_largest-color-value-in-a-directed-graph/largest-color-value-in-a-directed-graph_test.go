package main

import "testing"

func Test_largestPathValue(t *testing.T) {
	type args struct {
		colors string
		edges  [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"abaca", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}}}, 3},
		{"2", args{"a", [][]int{{0, 0}}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestPathValue(tt.args.colors, tt.args.edges); got != tt.want {
				t.Errorf("largestPathValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

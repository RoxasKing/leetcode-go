package main

import "testing"

func Test_maximumRequests(t *testing.T) {
	type args struct {
		n        int
		requests [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5, [][]int{{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4}}}, 5},
		{"2", args{3, [][]int{{0, 0}, {1, 2}, {2, 1}}}, 3},
		{"3", args{4, [][]int{{0, 3}, {3, 1}, {1, 2}, {2, 0}}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumRequests(tt.args.n, tt.args.requests); got != tt.want {
				t.Errorf("maximumRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}

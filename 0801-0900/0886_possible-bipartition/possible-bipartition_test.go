package main

import "testing"

func Test_possibleBipartition(t *testing.T) {
	type args struct {
		n        int
		dislikes [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{4, [][]int{{1, 2}, {1, 3}, {2, 4}}}, true},
		{"2", args{3, [][]int{{1, 2}, {1, 3}, {2, 3}}}, false},
		{"3", args{5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleBipartition(tt.args.n, tt.args.dislikes); got != tt.want {
				t.Errorf("possibleBipartition() = %v, want %v", got, tt.want)
			}
		})
	}
}

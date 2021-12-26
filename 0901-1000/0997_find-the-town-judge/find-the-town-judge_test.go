package main

import "testing"

func Test_findJudge(t *testing.T) {
	type args struct {
		n     int
		trust [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{2, [][]int{{1, 2}}}, 2},
		{"2", args{3, [][]int{{1, 3}, {2, 3}}}, 3},
		{"3", args{3, [][]int{{1, 3}, {2, 3}, {3, 1}}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findJudge(tt.args.n, tt.args.trust); got != tt.want {
				t.Errorf("findJudge() = %v, want %v", got, tt.want)
			}
		})
	}
}

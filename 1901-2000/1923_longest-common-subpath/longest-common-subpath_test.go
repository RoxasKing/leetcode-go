package main

import "testing"

func Test_longestCommonSubpath(t *testing.T) {
	type args struct {
		n     int
		paths [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5, [][]int{{0, 1, 0, 1, 0, 1, 0, 1, 0}, {0, 1, 3, 0, 1, 4, 0, 1, 0}}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubpath(tt.args.n, tt.args.paths); got != tt.want {
				t.Errorf("longestCommonSubpath() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_maxSumSubmatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{1, 0, 1}, {0, -2, 3}}, 2}, 2},
		{"2", args{[][]int{{2, 2, -1}}, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumSubmatrix(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("maxSumSubmatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

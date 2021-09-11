package main

import "testing"

func Test_orderOfLargestPlusSign(t *testing.T) {
	type args struct {
		n     int
		mines [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5, [][]int{{4, 2}}}, 2},
		{"2", args{1, [][]int{{0, 0}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orderOfLargestPlusSign(tt.args.n, tt.args.mines); got != tt.want {
				t.Errorf("orderOfLargestPlusSign() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		n    int
		cuts []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{7, []int{1, 3, 4, 5}}, 16},
		{"2", args{9, []int{5, 6, 1, 4, 2}}, 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.n, tt.args.cuts); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

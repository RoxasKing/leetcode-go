package main

import "testing"

func Test_shortestSubarray(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1}, 1}, 1},
		{"2", args{[]int{1, 2}, 4}, -1},
		{"3", args{[]int{2, -1, 2}, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestSubarray(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("shortestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

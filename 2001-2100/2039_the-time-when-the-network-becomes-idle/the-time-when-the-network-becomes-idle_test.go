package main

import "testing"

func Test_networkBecomesIdle(t *testing.T) {
	type args struct {
		edges    [][]int
		patience []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{0, 1}, {1, 2}}, []int{0, 2, 1}}, 8},
		{"2", args{[][]int{{0, 1}, {0, 2}, {1, 2}}, []int{0, 10, 10}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := networkBecomesIdle(tt.args.edges, tt.args.patience); got != tt.want {
				t.Errorf("networkBecomesIdle() = %v, want %v", got, tt.want)
			}
		})
	}
}

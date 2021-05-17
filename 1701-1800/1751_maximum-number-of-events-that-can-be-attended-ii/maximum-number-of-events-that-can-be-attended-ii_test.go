package main

import "testing"

func Test_maxValue(t *testing.T) {
	type args struct {
		events [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 1}}, 2}, 7},
		{"2", args{[][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 10}}, 2}, 10},
		{"3", args{[][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}, {4, 4, 4}}, 3}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValue(tt.args.events, tt.args.k); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

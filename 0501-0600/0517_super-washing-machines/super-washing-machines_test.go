package main

import "testing"

func Test_findMinMoves(t *testing.T) {
	type args struct {
		machines []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 0, 5}}, 3},
		{"2", args{[]int{0, 3, 0}}, 2},
		{"3", args{[]int{0, 2, 0}}, -1},
		{"4", args{[]int{0, 0, 11, 5}}, 8},
		{"5", args{[]int{100000, 0, 100000, 0, 100000, 0, 100000, 0, 100000, 0, 100000, 0}}, 50000},
		{"6", args{[]int{0, 0, 10, 0, 0, 0, 10, 0, 0, 0}}, 8},
		{"7", args{[]int{0, 0, 11, 0, 11, 0, 0, 2}}, 8},
		{"8", args{[]int{0, 0, 4, 0, 5, 0, 5, 0, 4}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinMoves(tt.args.machines); got != tt.want {
				t.Errorf("findMinMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

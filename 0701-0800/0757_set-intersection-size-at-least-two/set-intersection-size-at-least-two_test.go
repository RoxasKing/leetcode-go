package main

import "testing"

func Test_intersectionSizeTwo(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}}}, 3},
		{"2", args{[][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersectionSizeTwo(tt.args.intervals); got != tt.want {
				t.Errorf("intersectionSizeTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

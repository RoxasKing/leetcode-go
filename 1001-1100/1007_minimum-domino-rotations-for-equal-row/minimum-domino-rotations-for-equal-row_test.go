package main

import "testing"

func Test_minDominoRotations(t *testing.T) {
	type args struct {
		tops    []int
		bottoms []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 2}}, 2},
		{"2", args{[]int{3, 5, 1, 2, 3}, []int{3, 6, 3, 3, 4}}, -1},
		{"3", args{[]int{1, 1, 1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1, 1, 1}}, 0},
		{"4", args{[]int{1, 2, 1, 1, 1, 2, 2, 2}, []int{2, 1, 2, 2, 2, 2, 2, 2}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDominoRotations(tt.args.tops, tt.args.bottoms); got != tt.want {
				t.Errorf("minDominoRotations() = %v, want %v", got, tt.want)
			}
		})
	}
}

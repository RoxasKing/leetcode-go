package main

import "testing"

func Test_minCostToMoveChips(t *testing.T) {
	type args struct {
		position []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3}}, 1},
		{"2", args{[]int{2, 2, 2, 3, 3}}, 2},
		{"3", args{[]int{1, 1000000000}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostToMoveChips(tt.args.position); got != tt.want {
				t.Errorf("minCostToMoveChips() = %v, want %v", got, tt.want)
			}
		})
	}
}

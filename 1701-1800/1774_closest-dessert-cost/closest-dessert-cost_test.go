package main

import "testing"

func Test_closestCost(t *testing.T) {
	type args struct {
		baseCosts    []int
		toppingCosts []int
		target       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				[]int{1, 7},
				[]int{3, 4},
				10,
			},
			10,
		},
		{
			"2",
			args{
				[]int{2, 3},
				[]int{4, 5, 100},
				18,
			},
			17,
		},
		{
			"3",
			args{
				[]int{3, 10},
				[]int{2, 5},
				9,
			},
			8,
		},
		{
			"4",
			args{
				[]int{10},
				[]int{1},
				1,
			},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestCost(tt.args.baseCosts, tt.args.toppingCosts, tt.args.target); got != tt.want {
				t.Errorf("closestCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

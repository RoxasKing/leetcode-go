package main

import "testing"

func Test_profitableSchemes(t *testing.T) {
	type args struct {
		n         int
		minProfit int
		group     []int
		profit    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5, 3, []int{2, 2}, []int{2, 3}}, 2},
		{"2", args{10, 5, []int{2, 3, 5}, []int{6, 7, 8}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := profitableSchemes(tt.args.n, tt.args.minProfit, tt.args.group, tt.args.profit); got != tt.want {
				t.Errorf("profitableSchemes() = %v, want %v", got, tt.want)
			}
		})
	}
}

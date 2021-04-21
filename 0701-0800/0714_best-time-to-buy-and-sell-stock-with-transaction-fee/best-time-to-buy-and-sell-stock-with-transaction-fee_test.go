package main

import "testing"

func Test_maxProfitVI(t *testing.T) {
	type args struct {
		prices []int
		fee    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 3, 2, 8, 4, 9}, 2}, 8},
		{"2", args{[]int{1, 3, 7, 5, 10, 3}, 3}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices, tt.args.fee); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

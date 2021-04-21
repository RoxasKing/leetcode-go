package main

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		k      int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{2, []int{2, 4, 1}}, 2},
		{"2", args{2, []int{3, 2, 6, 5, 0, 3}}, 7},
		{"3", args{2, []int{3, 3, 5, 0, 0, 3, 1, 4}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

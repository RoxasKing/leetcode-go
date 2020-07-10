package leetcode

import (
	"testing"
)

func Test_maxProfitV(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 0, 2}}, 3},
		{"2", args{[]int{1, 2, 4}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitV(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfitV2(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 0, 2}}, 3},
		{"2", args{[]int{1, 2, 4}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitV2(tt.args.prices); got != tt.want {
				t.Errorf("maxProfitIII2() = %v, want %v", got, tt.want)
			}
		})
	}
}

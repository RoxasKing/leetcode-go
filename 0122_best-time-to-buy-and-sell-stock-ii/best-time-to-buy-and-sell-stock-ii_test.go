package main

import (
	"testing"
)

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{7, 1, 5, 3, 6, 4}}, 7},
		{"2", args{[]int{7, 1, 5, 4, 6, 4}}, 6},
		{"3", args{[]int{7, 1, 5, 5, 6, 4}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit2(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{7, 1, 5, 3, 6, 4}}, 7},
		{"2", args{[]int{7, 1, 5, 4, 6, 4}}, 6},
		{"3", args{[]int{7, 1, 5, 5, 6, 4}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit2(tt.args.prices); got != tt.want {
				t.Errorf("maxProfitII2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit3(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{7, 1, 5, 3, 6, 4}}, 7},
		{"2", args{[]int{7, 1, 5, 4, 6, 4}}, 6},
		{"3", args{[]int{7, 1, 5, 5, 6, 4}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit3(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit3() = %v, want %v", got, tt.want)
			}
		})
	}
}

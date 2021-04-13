package main

import (
	"testing"
)

func Test_purchasePlans(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 5, 3, 5}, 6}, 1},
		{"2", args{[]int{2, 2, 1, 9}, 10}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := purchasePlans(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("purchasePlans() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_purchasePlans2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 5, 3, 5}, 6}, 1},
		{"2", args{[]int{2, 2, 1, 9}, 10}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := purchasePlans2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("purchasePlans2() = %v, want %v", got, tt.want)
			}
		})
	}
}

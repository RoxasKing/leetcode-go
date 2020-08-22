package main

import (
	"testing"
)

func Test_judgePoint24(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{4, 1, 8, 7}}, true},
		{"2", args{[]int{1, 2, 1, 2}}, false},
		{"3", args{[]int{1, 5, 9, 1}}, false},
		{"4", args{[]int{1, 9, 1, 2}}, true},
		{"5", args{[]int{4, 7, 9, 9}}, true},
		{"6", args{[]int{7, 6, 7, 5}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgePoint24(tt.args.nums); got != tt.want {
				t.Errorf("judgePoint24() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_judgePoint24_2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{4, 1, 8, 7}}, true},
		{"2", args{[]int{1, 2, 1, 2}}, false},
		{"3", args{[]int{1, 5, 9, 1}}, false},
		{"4", args{[]int{1, 9, 1, 2}}, true},
		{"5", args{[]int{4, 7, 9, 9}}, true},
		{"6", args{[]int{7, 6, 7, 5}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgePoint24_2(tt.args.nums); got != tt.want {
				t.Errorf("judgePoint24_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

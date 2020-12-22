package main

import (
	"testing"
)

func Test_wiggleMaxLength(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 7, 4, 9, 2, 5}}, 6},
		{"2", args{[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}}, 7},
		{"3", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 2},
		{"4", args{[]int{1, 2, 3, 4, 6, 5, 6, 7, 8, 9}}, 4},
		{"5", args{[]int{1, 2}}, 2},
		{"6", args{[]int{2}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wiggleMaxLength(tt.args.nums); got != tt.want {
				t.Errorf("wiggleMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wiggleMaxLength2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 7, 4, 9, 2, 5}}, 6},
		{"2", args{[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}}, 7},
		{"3", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 2},
		{"4", args{[]int{1, 2, 3, 4, 6, 5, 6, 7, 8, 9}}, 4},
		{"5", args{[]int{1, 2}}, 2},
		{"6", args{[]int{2}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wiggleMaxLength2(tt.args.nums); got != tt.want {
				t.Errorf("wiggleMaxLength2() = %v, want %v", got, tt.want)
			}
		})
	}
}

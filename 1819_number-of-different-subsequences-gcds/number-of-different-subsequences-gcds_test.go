package main

import (
	"testing"
)

func Test_countDifferentSubsequenceGCDs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{6, 10, 3}}, 5},
		{"2", args{[]int{5, 15, 40, 5, 6}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDifferentSubsequenceGCDs(tt.args.nums); got != tt.want {
				t.Errorf("countDifferentSubsequenceGCDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

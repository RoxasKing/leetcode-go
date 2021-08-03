package main

import "testing"

func Test_countSpecialSubsequences(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{0, 1, 2, 2}}, 3},
		{"2", args{[]int{2, 2, 0, 0}}, 0},
		{"3", args{[]int{0, 1, 2, 0, 1, 2}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSpecialSubsequences(tt.args.nums); got != tt.want {
				t.Errorf("countSpecialSubsequences() = %v, want %v", got, tt.want)
			}
		})
	}
}

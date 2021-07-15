package main

import "testing"

func Test_waysToSplit(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 1, 1}}, 1},
		{"2", args{[]int{1, 2, 2, 2, 5, 0}}, 3},
		{"3", args{[]int{3, 2, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToSplit(tt.args.nums); got != tt.want {
				t.Errorf("waysToSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}

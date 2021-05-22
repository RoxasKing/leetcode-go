package main

import "testing"

func Test_maxUncrossedLines(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 4, 2}, []int{1, 2, 4}}, 2},
		{"2", args{[]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}}, 3},
		{"3", args{[]int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxUncrossedLines(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("maxUncrossedLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

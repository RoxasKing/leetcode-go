package main

import "testing"

func Test_minimumXORSum(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2}, []int{2, 3}}, 2},
		{"2", args{[]int{1, 0, 3}, []int{5, 3, 4}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumXORSum(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("minimumXORSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

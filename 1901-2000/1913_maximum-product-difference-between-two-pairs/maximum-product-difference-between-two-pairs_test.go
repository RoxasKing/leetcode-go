package main

import "testing"

func Test_maxProductDifference(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{5, 6, 2, 7, 4}}, 34},
		{"2", args{[]int{4, 2, 5, 9, 7, 4, 8}}, 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProductDifference(tt.args.nums); got != tt.want {
				t.Errorf("maxProductDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

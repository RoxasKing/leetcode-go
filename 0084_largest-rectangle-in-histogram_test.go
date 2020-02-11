package leetcode

import "testing"

func Test_largestRectangleArea(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{6, 4, 5, 2, 4, 3, 9}}, 14},
		{"", args{[]int{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 6, 2, 3}}, 15},
		{"", args{[]int{2, 1, 5, 6, 2, 3}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

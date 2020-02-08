package leetcode

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{4, 1, 2, 3}}, 5},
		{"", args{[]int{2, 1}}, 3},
		{"", args{[]int{1, 2, 0}}, 3},
		{"", args{[]int{3, 4, -1, 1}}, 2},
		{"", args{[]int{7, 8, 9, 11, 12}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}

package leetcode

import "testing"

func Test_removeDuplicates0080(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{1, 1, 1, 2, 2, 3}}, 5},
		{"", args{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates0080(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates0080() = %v, want %v", got, tt.want)
			}
		})
	}
}
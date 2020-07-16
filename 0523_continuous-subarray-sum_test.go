package leetcode

import "testing"

func Test_checkSubarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{23, 2, 4, 6, 7}, 6}, true},
		{"2", args{[]int{23, 2, 6, 4, 7}, 6}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSubarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("checkSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}

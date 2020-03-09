package leetcode

import "testing"

func Test_maxProfit0121(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{7, 1, 5, 3, 6, 4}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit0121(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit0121() = %v, want %v", got, tt.want)
			}
		})
	}
}

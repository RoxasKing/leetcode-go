package leetcode

import "testing"

func Test_maxProfitII(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{7, 1, 5, 3, 6, 4}}, 7},
		{"2", args{[]int{7, 1, 5, 4, 6, 4}}, 6},
		{"3", args{[]int{7, 1, 5, 5, 6, 4}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitII(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

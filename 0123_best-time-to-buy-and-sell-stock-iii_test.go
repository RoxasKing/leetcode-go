package leetcode

import "testing"

func Test_maxProfitIII(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{3, 3, 5, 0, 0, 3, 1, 4}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitIII(tt.args.prices); got != tt.want {
				t.Errorf("maxProfitIII() = %v, want %v", got, tt.want)
			}
		})
	}
}

package leetcode

import "testing"

func Test_maxProfitIV(t *testing.T) {
	type args struct {
		k      int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{2, []int{2, 4, 1}}, 2},
		{"2", args{3, []int{3, 2, 6, 5, 0, 3}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitIV(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfitIV() = %v, want %v", got, tt.want)
			}
		})
	}
}

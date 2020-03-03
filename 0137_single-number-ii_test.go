package leetcode

import "testing"

func Test_singleNumber0137(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{2, 2, 2, 3}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber0137(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber0137() = %v, want %v", got, tt.want)
			}
		})
	}
}

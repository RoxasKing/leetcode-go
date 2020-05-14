package leetcode

import "testing"

func Test_singleNumberII(t *testing.T) {
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
			if got := singleNumberII(tt.args.nums); got != tt.want {
				t.Errorf("singleNumberII() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"testing"
)

func Test_exchange(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 2, 3, 4}}, []int{1, 3, 2, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := exchange(tt.args.nums)
			l, r := 0, len(got)-1
			for l < len(got) && got[l]&1 == 1 {
				l++
			}
			for r >= 0 && got[r]&1 == 0 {
				r--
			}
			if l <= r {
				t.Errorf("exchange() = %v, want %v", got, tt.want)
			}
		})
	}
}

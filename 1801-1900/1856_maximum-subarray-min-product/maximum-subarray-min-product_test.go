package main

import "testing"

func Test_maxSumMinProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 2}}, 14},
		{"2", args{[]int{2, 3, 3, 1, 2}}, 18},
		{"3", args{[]int{3, 1, 5, 6, 4, 2}}, 60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumMinProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxSumMinProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

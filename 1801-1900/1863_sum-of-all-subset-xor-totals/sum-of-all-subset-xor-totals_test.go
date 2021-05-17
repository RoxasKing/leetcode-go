package main

import "testing"

func Test_subsetXORSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 3}}, 6},
		{"2", args{[]int{5, 1, 6}}, 28},
		{"3", args{[]int{3, 4, 5, 6, 7, 8}}, 480},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsetXORSum(tt.args.nums); got != tt.want {
				t.Errorf("subsetXORSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

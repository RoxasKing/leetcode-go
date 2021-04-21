package main

import "testing"

func Test_arrayPairSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 4, 3, 2}}, 4},
		{"2", args{[]int{6, 2, 6, 5, 1, 2}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayPairSum(tt.args.nums); got != tt.want {
				t.Errorf("arrayPairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_sumOfFlooredPairs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 5, 9}}, 10},
		{"2", args{[]int{7, 7, 7, 7, 7, 7, 7}}, 49},
		{"3", args{[]int{4, 3, 4, 3, 5}}, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfFlooredPairs(tt.args.nums); got != tt.want {
				t.Errorf("sumOfFlooredPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

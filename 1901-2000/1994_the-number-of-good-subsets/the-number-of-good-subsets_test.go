package main

import "testing"

func Test_numberOfGoodSubsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 4}}, 6},
		{"2", args{[]int{4, 2, 3, 15}}, 5},
		{"3", args{[]int{12, 3}}, 1},
		{"4", args{[]int{5, 10, 1, 26, 24, 21, 24, 23, 11, 12, 27, 4, 17, 16, 2, 6, 1, 1, 6, 8, 13, 30, 24, 20, 2, 19}}, 5368},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfGoodSubsets(tt.args.nums); got != tt.want {
				t.Errorf("numberOfGoodSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

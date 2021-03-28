package main

import "testing"

func Test_findMaximumXOR(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{3, 10, 5, 25, 2, 8}}, 28},
		{"2", args{[]int{0}}, 0},
		{"3", args{[]int{2, 4}}, 6},
		{"4", args{[]int{8, 10, 2}}, 10},
		{"5", args{[]int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}}, 127},
		{"6", args{[]int{32, 18, 33, 42, 29, 20, 26, 36, 15, 46}}, 62},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaximumXOR(tt.args.nums); got != tt.want {
				t.Errorf("findMaximumXOR() = %v, want %v", got, tt.want)
			}
		})
	}
}

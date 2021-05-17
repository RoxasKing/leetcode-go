package main

import "testing"

func Test_smallestDivisor(t *testing.T) {
	type args struct {
		nums      []int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 5, 9}, 6}, 5},
		{"2", args{[]int{44, 22, 33, 11, 1}, 5}, 44},
		{"3", args{[]int{21212, 10101, 12121}, 1000000}, 1},
		{"4", args{[]int{2, 3, 5, 7, 11}, 11}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestDivisor(tt.args.nums, tt.args.threshold); got != tt.want {
				t.Errorf("smallestDivisor() = %v, want %v", got, tt.want)
			}
		})
	}
}

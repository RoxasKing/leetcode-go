package main

import "testing"

func Test_maxEqualFreq(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 2, 1, 1, 5, 3, 3, 5}}, 7},
		{"2", args{[]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5}}, 13},
		{"3", args{[]int{1, 2}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEqualFreq(tt.args.nums); got != tt.want {
				t.Errorf("maxEqualFreq() = %v, want %v", got, tt.want)
			}
		})
	}
}

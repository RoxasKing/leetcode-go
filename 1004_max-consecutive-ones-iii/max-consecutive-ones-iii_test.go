package main

import "testing"

func Test_longestOnes(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2}, 6},
		{"2", args{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3}, 10},
		{"3", args{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 0}, 4},
		{"4", args{[]int{1}, 3}, 1},
		{"5", args{[]int{0}, 3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestOnes(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("longestOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}

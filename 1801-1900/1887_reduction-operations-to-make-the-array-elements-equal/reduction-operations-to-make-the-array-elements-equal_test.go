package main

import "testing"

func Test_reductionOperations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{5, 1, 3}}, 3},
		{"2", args{[]int{1, 1, 1}}, 0},
		{"3", args{[]int{1, 1, 2, 2, 3}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reductionOperations(tt.args.nums); got != tt.want {
				t.Errorf("reductionOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

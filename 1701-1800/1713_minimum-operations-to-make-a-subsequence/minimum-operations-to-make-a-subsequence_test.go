package main

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		target []int
		arr    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{5, 1, 3}, []int{9, 4, 2, 3, 4}}, 2},
		{"2", args{[]int{6, 4, 8, 1, 3, 2}, []int{4, 7, 6, 2, 3, 8, 6, 1}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.target, tt.args.arr); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

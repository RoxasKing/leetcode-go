package main

import (
	"testing"
)

func Test_minIncrementForUnique(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 2}}, 1},
		{"2", args{[]int{3, 2, 1, 2, 1, 7}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minIncrementForUnique(tt.args.A); got != tt.want {
				t.Errorf("minIncrementForUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

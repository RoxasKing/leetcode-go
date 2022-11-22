package main

import "testing"

func Test_isIdealPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{1, 0, 2}}, true},
		{"2", args{[]int{1, 2, 0}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIdealPermutation(tt.args.nums); got != tt.want {
				t.Errorf("isIdealPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

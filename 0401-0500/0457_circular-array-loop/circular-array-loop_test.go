package main

import "testing"

func Test_circularArrayLoop(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{2, -1, 1, 2, 2}}, true},
		{"2", args{[]int{-1, 2}}, false},
		{"3", args{[]int{-2, 1, -1, -2, -2}}, false},
		{"4", args{[]int{-1, -2, -3, -4, -5}}, false},
		{"5", args{[]int{3, 1, 2}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := circularArrayLoop(tt.args.nums); got != tt.want {
				t.Errorf("circularArrayLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}

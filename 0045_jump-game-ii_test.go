package main

import (
	"testing"
)

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{2, 3, 1, 1, 4}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jump_2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{2, 3, 1, 1, 4}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump2(tt.args.nums); got != tt.want {
				t.Errorf("jump_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

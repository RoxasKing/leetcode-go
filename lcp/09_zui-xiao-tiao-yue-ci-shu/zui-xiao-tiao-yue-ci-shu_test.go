package main

import (
	"testing"
)

func Test_minJump(t *testing.T) {
	type args struct {
		jump []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 5, 1, 1, 1, 1}}, 3},
		{"2", args{[]int{2, 2, 1, 1, 1, 1}}, 5},
		{"3", args{[]int{2, 4, 1, 1, 1, 1}}, 4},
		{"4", args{[]int{2, 5, 4, 1, 1, 1}}, 2},
		{"5", args{[]int{2, 7, 1, 1, 1, 1, 1, 1, 1}}, 4},
		{"6", args{[]int{2, 7, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 1, 1, 2, 1, 6, 1, 1, 1, 1, 1}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minJump(tt.args.jump); got != tt.want {
				t.Errorf("minJump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minJump2(t *testing.T) {
	type args struct {
		jump []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 5, 1, 1, 1, 1}}, 3},
		{"2", args{[]int{2, 2, 1, 1, 1, 1}}, 5},
		{"3", args{[]int{2, 4, 1, 1, 1, 1}}, 4},
		{"4", args{[]int{2, 5, 4, 1, 1, 1}}, 2},
		{"5", args{[]int{2, 7, 1, 1, 1, 1, 1, 1, 1}}, 4},
		{"6", args{[]int{2, 7, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 1, 1, 2, 1, 6, 1, 1, 1, 1, 1}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minJump2(tt.args.jump); got != tt.want {
				t.Errorf("minJump2() = %v, want %v", got, tt.want)
			}
		})
	}
}

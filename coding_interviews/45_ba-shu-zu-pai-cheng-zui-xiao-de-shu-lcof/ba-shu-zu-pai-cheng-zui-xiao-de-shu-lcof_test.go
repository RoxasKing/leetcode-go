package main

import (
	"testing"
)

func Test_minNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{[]int{10, 2}}, "102"},
		{"2", args{[]int{3, 30, 34, 5, 9}}, "3033459"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumber(tt.args.nums); got != tt.want {
				t.Errorf("minNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minNumber2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{[]int{10, 2}}, "102"},
		{"2", args{[]int{3, 30, 34, 5, 9}}, "3033459"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumber2(tt.args.nums); got != tt.want {
				t.Errorf("minNumber2() = %v, want %v", got, tt.want)
			}
		})
	}
}

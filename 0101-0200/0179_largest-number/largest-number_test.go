package main

import "testing"

func Test_largestNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{[]int{10, 2}}, "210"},
		{"2", args{[]int{3, 30, 34, 5, 9}}, "9534330"},
		{"3", args{[]int{0, 1, 2, 3}}, "3210"},
		{"4", args{[]int{0, 0, 0, 0}}, "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.args.nums); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_check(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{3, 4, 5, 1, 2}}, true},
		{"2", args{[]int{2, 1, 3, 4}}, false},
		{"3", args{[]int{1, 2, 3}}, true},
		{"4", args{[]int{1, 1, 1}}, true},
		{"5", args{[]int{2, 1}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.nums); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_checkPossibility(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{4, 2, 3}}, true},
		{"2", args{[]int{4, 2, 1}}, false},
		{"3", args{[]int{1, 2, 4, 3, 4, 5}}, true},
		{"4", args{[]int{1, 2, 4, 1, 4, 5}}, true},
		{"5", args{[]int{1, 2, 4, 1, 3, 5}}, false},
		{"6", args{[]int{1, 2, 4, 2, 3, 5}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPossibility(tt.args.nums); got != tt.want {
				t.Errorf("checkPossibility() = %v, want %v", got, tt.want)
			}
		})
	}
}

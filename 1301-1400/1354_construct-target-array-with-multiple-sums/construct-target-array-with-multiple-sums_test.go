package main

import "testing"

func Test_isPossible(t *testing.T) {
	type args struct {
		target []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{9, 3, 5}}, true},
		{"2", args{[]int{1, 1, 1, 2}}, false},
		{"3", args{[]int{8, 5}}, true},
		{"4", args{[]int{2}}, false},
		{"5", args{[]int{1, 1000000000}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPossible(tt.args.target); got != tt.want {
				t.Errorf("isPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}

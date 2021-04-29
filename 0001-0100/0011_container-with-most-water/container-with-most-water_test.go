package main

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
		{"2", args{[]int{1, 1}}, 1},
		{"3", args{[]int{4, 3, 2, 1, 4}}, 16},
		{"4", args{[]int{1, 2, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

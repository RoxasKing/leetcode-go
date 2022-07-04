package main

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		h              int
		w              int
		horizontalCuts []int
		verticalCuts   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5, 4, []int{1, 2, 4}, []int{1, 3}}, 4},
		{"2", args{5, 4, []int{3, 1}, []int{1}}, 6},
		{"3", args{5, 4, []int{3}, []int{3}}, 9},
		{"4", args{1000000000, 1000000000, []int{2}, []int{2}}, 81},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.h, tt.args.w, tt.args.horizontalCuts, tt.args.verticalCuts); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_rectangleArea(t *testing.T) {
	type args struct {
		rectangles [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{0, 0, 2, 2}, {1, 0, 2, 3}, {1, 0, 3, 1}}}, 6},
		{"2", args{[][]int{{0, 0, 1000000000, 1000000000}}}, 49},
		{"3", args{[][]int{{0, 0, 1, 1}, {2, 2, 3, 3}}}, 2},
		{"4", args{[][]int{{49, 40, 62, 100}, {11, 83, 31, 99}, {19, 39, 30, 99}}}, 1584},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rectangleArea(tt.args.rectangles); got != tt.want {
				t.Errorf("rectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

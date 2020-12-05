package main

import "testing"

func Test_largestPerimeter(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 1, 2}}, 5},
		{"2", args{[]int{1, 2, 1}}, 0},
		{"3", args{[]int{3, 2, 3, 4}}, 10},
		{"4", args{[]int{3, 6, 2, 3}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestPerimeter(tt.args.A); got != tt.want {
				t.Errorf("largestPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

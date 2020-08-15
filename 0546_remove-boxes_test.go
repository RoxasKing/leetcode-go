package leetcode

import (
	"testing"
)

func Test_removeBoxes(t *testing.T) {
	type args struct {
		boxes []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 3, 2, 2, 2, 3, 4, 3, 1}}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeBoxes(tt.args.boxes); got != tt.want {
				t.Errorf("removeBoxes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeBoxes2(t *testing.T) {
	type args struct {
		boxes []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 3, 2, 2, 2, 3, 4, 3, 1}}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeBoxes2(tt.args.boxes); got != tt.want {
				t.Errorf("removeBoxes2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package leetcode

import (
	"testing"
)

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{6, 5, 5}}, 5},
		{"", args{[]int{3, 2, 3}}, 3},
		{"", args{[]int{2, 2, 1, 1, 1, 2, 2}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_majorityElement2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{3, 3, 4}}, 3},
		{"", args{[]int{6, 5, 5}}, 5},
		{"", args{[]int{3, 2, 3}}, 3},
		{"", args{[]int{2, 2, 1, 1, 1, 2, 2}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement2(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_majorityElement3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{6, 5, 5}}, 5},
		{"", args{[]int{3, 2, 3}}, 3},
		{"", args{[]int{2, 2, 1, 1, 1, 2, 2}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement3(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement3() = %v, want %v", got, tt.want)
			}
		})
	}
}

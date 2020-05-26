package leetcode

import (
	"testing"
)

func Test_findDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 3, 4, 2, 2}}, 2},
		{"2", args{[]int{3, 1, 3, 4, 2}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDuplicate2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 3, 4, 2, 2}}, 2},
		{"2", args{[]int{3, 1, 3, 4, 2}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate2(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDuplicate3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 3, 4, 2, 2}}, 2},
		{"2", args{[]int{3, 1, 3, 4, 2}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate3(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate3() = %v, want %v", got, tt.want)
			}
		})
	}
}

package leetcode

import (
	"testing"
)

func Test_candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 87, 87, 87, 2, 1}}, 13},
		{"2", args{[]int{1, 3, 2, 2, 1}}, 7},
		{"3", args{[]int{1, 3, 2, 2, 2, 1}}, 8},
		{"4", args{[]int{1, 0, 2}}, 5},
		{"5", args{[]int{1, 2, 2}}, 4},
		{"6", args{[]int{1, 2, 1}}, 4},
		{"7", args{[]int{1, 4, 3, 3, 3, 1}}, 8},
		{"8", args{[]int{1, 4, 3, 2, 2, 1}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy(tt.args.ratings); got != tt.want {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_candy2(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 87, 87, 87, 2, 1}}, 13},
		{"2", args{[]int{1, 3, 2, 2, 1}}, 7},
		{"3", args{[]int{1, 3, 2, 2, 2, 1}}, 8},
		{"4", args{[]int{1, 0, 2}}, 5},
		{"5", args{[]int{1, 2, 2}}, 4},
		{"6", args{[]int{1, 2, 1}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy2(tt.args.ratings); got != tt.want {
				t.Errorf("candy2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_candy3(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 87, 87, 87, 2, 1}}, 13},
		{"2", args{[]int{1, 3, 2, 2, 1}}, 7},
		{"3", args{[]int{1, 3, 2, 2, 2, 1}}, 8},
		{"4", args{[]int{1, 0, 2}}, 5},
		{"5", args{[]int{1, 2, 2}}, 4},
		{"6", args{[]int{1, 2, 1}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy3(tt.args.ratings); got != tt.want {
				t.Errorf("candy3() = %v, want %v", got, tt.want)
			}
		})
	}
}

package leetcode

import (
	"testing"
)

func Test_hasGroupsSizeX(t *testing.T) {
	type args struct {
		deck []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{1, 2, 3, 4, 4, 3, 2, 1}}, true},
		{"2", args{[]int{1, 1, 1, 2, 2, 2, 3, 3}}, false},
		{"3", args{[]int{1}}, false},
		{"4", args{[]int{1, 1}}, true},
		{"5", args{[]int{1, 1, 2, 2, 2, 2}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasGroupsSizeX(tt.args.deck); got != tt.want {
				t.Errorf("hasGroupsSizeX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasGroupsSizeX2(t *testing.T) {
	type args struct {
		deck []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{1, 2, 3, 4, 4, 3, 2, 1}}, true},
		{"2", args{[]int{1, 1, 1, 2, 2, 2, 3, 3}}, false},
		{"3", args{[]int{1}}, false},
		{"4", args{[]int{1, 1}}, true},
		{"5", args{[]int{1, 1, 2, 2, 2, 2}}, true},
		{"6", args{[]int{1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3}}, false},
		{"7", args{[]int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasGroupsSizeX2(tt.args.deck); got != tt.want {
				t.Errorf("hasGroupsSizeX2() = %v, want %v", got, tt.want)
			}
		})
	}
}

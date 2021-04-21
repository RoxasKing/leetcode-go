package main

import (
	"testing"
)

func Test_canThreePartsEqualSum(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{[]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}}, true},
		{"", args{[]int{0, 2, 1, -6, 6, -7, 9, -1, 2, 0, 1}}, false},
		{"", args{[]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}}, true},
		{"", args{[]int{1, -1, 1, -1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canThreePartsEqualSum(tt.args.A); got != tt.want {
				t.Errorf("canThreePartsEqualSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

package leetcode

import (
	"testing"
)

func Test_canMeasureWater(t *testing.T) {
	type args struct {
		x int
		y int
		z int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{3, 5, 4}, true},
		{"", args{2, 6, 5}, false},
		{"", args{34, 5, 6}, true},
		{"", args{0, 0, 0}, true},
		{"", args{1, 2, 3}, true},
		{"", args{6, 9, 1}, false},
		{"", args{1, 1, 12}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMeasureWater(tt.args.x, tt.args.y, tt.args.z); got != tt.want {
				t.Errorf("canMeasureWater() = %v, want %v", got, tt.want)
			}
		})
	}
}

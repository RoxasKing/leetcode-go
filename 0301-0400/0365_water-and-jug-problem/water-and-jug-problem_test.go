package main

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
		{"1", args{3, 5, 4}, true},
		{"2", args{2, 6, 5}, false},
		{"3", args{1, 2, 3}, true},
		{"4", args{34, 5, 6}, true},
		{"5", args{0, 0, 0}, true},
		{"6", args{6, 9, 1}, false},
		{"7", args{1, 1, 12}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMeasureWater(tt.args.x, tt.args.y, tt.args.z); got != tt.want {
				t.Errorf("canMeasureWater() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"math"
	"testing"
)

func Test_maxAverageRatio(t *testing.T) {
	type args struct {
		classes       [][]int
		extraStudents int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1", args{[][]int{{1, 2}, {3, 5}, {2, 2}}, 2}, 0.78333},
		{"2", args{[][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}}, 4}, 0.53485},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAverageRatio(tt.args.classes, tt.args.extraStudents); math.Abs(got-tt.want) < 1e-9 {
				t.Errorf("maxAverageRatio() = %v, want %v", got, tt.want)
			}
		})
	}
}

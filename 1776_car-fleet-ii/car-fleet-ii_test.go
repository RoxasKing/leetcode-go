package main

import (
	"math"
	"testing"
)

func Test_getCollisionTimes(t *testing.T) {
	type args struct {
		cars [][]int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{"1", args{[][]int{{1, 2}, {2, 1}, {4, 3}, {7, 2}}}, []float64{1.00000, -1.00000, 3.00000, -1.00000}},
		{"2", args{[][]int{{3, 4}, {5, 4}, {6, 3}, {9, 1}}}, []float64{2.00000, 1.00000, 1.50000, -1.00000}},
		{
			"3",
			args{[][]int{{1, 4}, {4, 5}, {7, 1}, {13, 4}, {14, 3}, {15, 2}, {16, 5}, {19, 1}}},
			[]float64{2.00000, 0.75000, -1.00000, 1.00000, 1.00000, 4.00000, 0.75000, -1.00000},
		},
		{
			"4",
			args{[][]int{{1, 5}, {5, 5}, {6, 2}, {7, 2}, {9, 3}, {10, 4}, {12, 5}, {18, 2}, {20, 1}}},
			[]float64{1.66667, 0.33333, 14.0, 13.0, 5.5, 3.33333, 2.0, 2.0, -1.0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getCollisionTimes(tt.args.cars)
			for i := range got {
				if math.Abs(got[i]-tt.want[i]) > 1e-5 {
					t.Errorf("getCollisionTimes() = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}

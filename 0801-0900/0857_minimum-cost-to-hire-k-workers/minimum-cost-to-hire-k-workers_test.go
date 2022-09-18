package main

import (
	"math"
	"testing"
)

func Test_mincostToHireWorkers(t *testing.T) {
	type args struct {
		quality []int
		wage    []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1", args{[]int{10, 20, 5}, []int{70, 50, 30}, 2}, 105.00000},
		{"2", args{[]int{3, 1, 10, 10, 1}, []int{4, 8, 2, 2, 7}, 3}, 30.66667},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mincostToHireWorkers(tt.args.quality, tt.args.wage, tt.args.k); math.Abs(got-tt.want) > 1e-5 {
				t.Errorf("mincostToHireWorkers() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_medianSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{"1", args{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3}, []float64{1, -1, -1, 3, 5, 6}},
		{"2", args{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 4}, []float64{0.000, 1.000, 1.000, 4.000, 5.500}},
		{"3", args{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 5}, []float64{1.000, 3.000, 3.000, 5.000}},
		{"4", args{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 6}, []float64{2.000, 3.000, 4.000}},
		{"5", args{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 2}, []float64{2.000, 1.000, -2.000, 1.000, 4.000, 4.500, 6.500}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := medianSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("medianSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

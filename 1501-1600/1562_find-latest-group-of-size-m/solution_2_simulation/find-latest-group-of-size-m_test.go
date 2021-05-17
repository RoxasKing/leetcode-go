package main

import "testing"

func Test_findLatestStep(t *testing.T) {
	type args struct {
		arr []int
		m   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{3, 5, 1, 2, 4}, 1}, 4},
		{"2", args{[]int{3, 1, 5, 4, 2}, 2}, -1},
		{"3", args{[]int{1}, 1}, 1},
		{"4", args{[]int{2, 1}, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLatestStep(tt.args.arr, tt.args.m); got != tt.want {
				t.Errorf("findLatestStep() = %v, want %v", got, tt.want)
			}
		})
	}
}

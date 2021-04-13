package main

import (
	"testing"
)

func Test_minTime(t *testing.T) {
	type args struct {
		time []int
		m    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 3}, 2}, 3},
		{"2", args{[]int{999, 999, 999}, 4}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTime(tt.args.time, tt.args.m); got != tt.want {
				t.Errorf("minTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

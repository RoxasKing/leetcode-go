package main

import (
	"testing"
)

func Test_minSteps(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{3}, 3},
		{"2", args{25}, 10},
		{"3", args{0}, 0},
		{"4", args{1}, 0},
		{"5", args{199}, 199},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSteps(tt.args.n); got != tt.want {
				t.Errorf("minSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minSteps2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{3}, 3},
		{"2", args{25}, 10},
		{"3", args{0}, 0},
		{"4", args{1}, 0},
		{"5", args{199}, 199},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSteps2(tt.args.n); got != tt.want {
				t.Errorf("minSteps2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_colorTheGrid(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1, 1}, 3},
		{"2", args{1, 2}, 6},
		{"3", args{5, 5}, 580986},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := colorTheGrid(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("colorTheGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

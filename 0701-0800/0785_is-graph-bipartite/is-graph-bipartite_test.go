package main

import (
	"testing"
)

func Test_isBipartite(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}}, true},
		{"2", args{[][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBipartite(tt.args.graph); got != tt.want {
				t.Errorf("isBipartite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isBipartite2(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}}, true},
		{"2", args{[][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBipartite2(tt.args.graph); got != tt.want {
				t.Errorf("isBipartite2() = %v, want %v", got, tt.want)
			}
		})
	}
}

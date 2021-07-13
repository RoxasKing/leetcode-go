package main

import "testing"

func Test_nearestExit(t *testing.T) {
	type args struct {
		maze     [][]byte
		entrance []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nearestExit(tt.args.maze, tt.args.entrance); got != tt.want {
				t.Errorf("nearestExit() = %v, want %v", got, tt.want)
			}
		})
	}
}

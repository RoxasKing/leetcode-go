package main

import "testing"

func Test_robot(t *testing.T) {
	type args struct {
		command   string
		obstacles [][]int
		x         int
		y         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"URR", [][]int{}, 3, 2}, true},
		{"2", args{"URR", [][]int{{2, 2}}, 3, 2}, false},
		{"3", args{"URR", [][]int{{4, 2}}, 3, 2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robot(tt.args.command, tt.args.obstacles, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("robot() = %v, want %v", got, tt.want)
			}
		})
	}
}

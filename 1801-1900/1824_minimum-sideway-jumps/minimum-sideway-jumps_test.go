package main

import "testing"

func Test_minSideJumps(t *testing.T) {
	type args struct {
		obstacles []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{0, 1, 2, 3, 0}}, 2},
		{"2", args{[]int{0, 1, 1, 3, 3, 0}}, 0},
		{"3", args{[]int{0, 2, 1, 0, 3, 0}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSideJumps(tt.args.obstacles); got != tt.want {
				t.Errorf("minSideJumps() = %v, want %v", got, tt.want)
			}
		})
	}
}

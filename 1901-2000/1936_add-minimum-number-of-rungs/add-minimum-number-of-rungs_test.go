package main

import "testing"

func Test_addRungs(t *testing.T) {
	type args struct {
		rungs []int
		dist  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 3, 5, 10}, 2}, 2},
		{"2", args{[]int{3, 6, 8, 10}, 3}, 0},
		{"3", args{[]int{3, 4, 6, 7}, 2}, 1},
		{"4", args{[]int{5}, 10}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addRungs(tt.args.rungs, tt.args.dist); got != tt.want {
				t.Errorf("addRungs() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_waysToBuildRooms(t *testing.T) {
	type args struct {
		prevRoom []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{-1, 0, 1}}, 1},
		{"2", args{[]int{-1, 0, 0, 1, 2}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToBuildRooms(tt.args.prevRoom); got != tt.want {
				t.Errorf("waysToBuildRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}

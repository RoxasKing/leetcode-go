package main

import "testing"

func Test_findInMountainArray(t *testing.T) {
	type args struct {
		target      int
		mountainArr *MountainArray
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{3, Constructor([]int{1, 2, 3, 4, 5, 3, 1})}, 2},
		{"2", args{3, Constructor([]int{0, 1, 2, 4, 2, 1})}, -1},
		{"3", args{0, Constructor([]int{3, 5, 3, 2, 0})}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findInMountainArray(tt.args.target, tt.args.mountainArr); got != tt.want {
				t.Errorf("findInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

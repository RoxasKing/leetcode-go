package main

import (
	"testing"
)

func Test_maxHappyGroups(t *testing.T) {
	type args struct {
		batchSize int
		groups    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{3, []int{1, 2, 3, 4, 5, 6}}, 4},
		{"2", args{4, []int{1, 3, 2, 5, 2, 2, 1, 6}}, 4},
		{"3", args{3, []int{369821235, 311690424, 74641571, 179819879, 171396603, 274036220}}, 4},
		{"4", args{7, []int{2, 7, 5, 2, 3, 2, 6, 5, 3, 6, 2, 3, 7, 2, 2, 5, 4, 6, 6, 4, 7, 5, 6, 1, 6, 2, 6, 6, 2, 5}}, 15},
		{"5", args{9, []int{3, 3, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxHappyGroups(tt.args.batchSize, tt.args.groups); got != tt.want {
				t.Errorf("maxHappyGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

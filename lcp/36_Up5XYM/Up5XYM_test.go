package main

import "testing"

func Test_maxGroupNumber(t *testing.T) {
	type args struct {
		tiles []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 2, 2, 3, 4}}, 1},
		{"2", args{[]int{2, 2, 2, 3, 4, 1, 3}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxGroupNumber(tt.args.tiles); got != tt.want {
				t.Errorf("maxGroupNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

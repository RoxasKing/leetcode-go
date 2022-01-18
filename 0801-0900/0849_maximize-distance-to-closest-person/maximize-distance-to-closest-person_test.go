package main

import "testing"

func Test_maxDistToClosest(t *testing.T) {
	type args struct {
		seats []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 0, 0, 0, 1, 0, 1}}, 2},
		{"2", args{[]int{1, 0, 0, 0}}, 3},
		{"4", args{[]int{0, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistToClosest(tt.args.seats); got != tt.want {
				t.Errorf("maxDistToClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}

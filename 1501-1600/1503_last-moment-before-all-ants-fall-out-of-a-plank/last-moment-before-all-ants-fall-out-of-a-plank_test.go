package main

import "testing"

func Test_getLastMoment(t *testing.T) {
	type args struct {
		n     int
		left  []int
		right []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{4, []int{4, 3}, []int{0, 1}}, 4},
		{"2", args{7, []int{}, []int{0, 1, 2, 3, 4, 5, 6, 7}}, 7},
		{"3", args{7, []int{0, 1, 2, 3, 4, 5, 6, 7}, []int{}}, 7},
		{"4", args{9, []int{5}, []int{4}}, 5},
		{"5", args{6, []int{6}, []int{0}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLastMoment(tt.args.n, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("getLastMoment() = %v, want %v", got, tt.want)
			}
		})
	}
}

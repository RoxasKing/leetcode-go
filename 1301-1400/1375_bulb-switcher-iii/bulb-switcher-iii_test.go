package main

import "testing"

func Test_numTimesAllBlue(t *testing.T) {
	type args struct {
		light []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 1, 3, 5, 4}}, 3},
		{"2", args{[]int{3, 2, 4, 1, 5}}, 2},
		{"3", args{[]int{4, 1, 2, 3}}, 1},
		{"4", args{[]int{2, 1, 4, 3, 6, 5}}, 3},
		{"5", args{[]int{1, 2, 3, 4, 5, 6}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTimesAllBlue(tt.args.light); got != tt.want {
				t.Errorf("numTimesAllBlue() = %v, want %v", got, tt.want)
			}
		})
	}
}

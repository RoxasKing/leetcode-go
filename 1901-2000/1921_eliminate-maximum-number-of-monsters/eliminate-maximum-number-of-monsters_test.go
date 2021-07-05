package main

import "testing"

func Test_eliminateMaximum(t *testing.T) {
	type args struct {
		dist  []int
		speed []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 3, 4}, []int{1, 1, 1}}, 3},
		{"2", args{[]int{1, 1, 2, 3}, []int{1, 1, 1, 1}}, 1},
		{"3", args{[]int{3, 2, 4}, []int{5, 3, 2}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eliminateMaximum(tt.args.dist, tt.args.speed); got != tt.want {
				t.Errorf("eliminateMaximum() = %v, want %v", got, tt.want)
			}
		})
	}
}

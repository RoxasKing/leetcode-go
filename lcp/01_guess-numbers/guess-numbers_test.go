package main

import "testing"

func Test_game(t *testing.T) {
	type args struct {
		guess  []int
		answer []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3}, []int{1, 2, 3}}, 3},
		{"2", args{[]int{2, 2, 3}, []int{3, 2, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := game(tt.args.guess, tt.args.answer); got != tt.want {
				t.Errorf("game() = %v, want %v", got, tt.want)
			}
		})
	}
}

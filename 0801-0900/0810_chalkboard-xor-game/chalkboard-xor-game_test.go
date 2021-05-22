package main

import "testing"

func Test_xorGame(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{1, 1, 2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xorGame(tt.args.nums); got != tt.want {
				t.Errorf("xorGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_magicTower(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{100, 100, 100, -250, -60, -140, -50, -50, 100, 150}}, 1},
		{"2", args{[]int{-200, -300, 400, 0}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := magicTower(tt.args.nums); got != tt.want {
				t.Errorf("magicTower() = %v, want %v", got, tt.want)
			}
		})
	}
}

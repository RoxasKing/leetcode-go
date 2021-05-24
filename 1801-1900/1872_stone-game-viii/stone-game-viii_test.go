package main

import "testing"

func Test_stoneGameVIII(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{-1, 2, -3, 4, -5}}, 5},
		{"2", args{[]int{7, -6, 5, 10, 5, -2, -6}}, 13},
		{"3", args{[]int{-10, -12}}, -22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGameVIII(tt.args.stones); got != tt.want {
				t.Errorf("stoneGameVIII() = %v, want %v", got, tt.want)
			}
		})
	}
}

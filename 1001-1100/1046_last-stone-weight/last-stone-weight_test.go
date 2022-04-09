package main

import "testing"

func Test_lastStoneWeight(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 7, 4, 1, 8, 1}}, 1},
		{"2", args{[]int{1}}, 1},
		{"3", args{[]int{2, 2}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastStoneWeight(tt.args.stones); got != tt.want {
				t.Errorf("lastStoneWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

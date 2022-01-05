package main

import "testing"

func Test_numPairsDivisibleBy60(t *testing.T) {
	type args struct {
		time []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{30, 20, 150, 100, 40}}, 3},
		{"2", args{[]int{60, 60, 60}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numPairsDivisibleBy60(tt.args.time); got != tt.want {
				t.Errorf("numPairsDivisibleBy60() = %v, want %v", got, tt.want)
			}
		})
	}
}

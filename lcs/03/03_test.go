package main

import "testing"

func Test_largestArea(t *testing.T) {
	type args struct {
		grid []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]string{"110", "231", "221"}}, 1},
		{"2", args{[]string{"11111100000", "21243101111", "21224101221", "11111101111"}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestArea(tt.args.grid); got != tt.want {
				t.Errorf("largestArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

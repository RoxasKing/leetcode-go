package main

import "testing"

func Test_maximumScore(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{2, 4, 6}, 6},
		{"2", args{4, 4, 6}, 7},
		{"3", args{1, 8, 8}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumScore(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("maximumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

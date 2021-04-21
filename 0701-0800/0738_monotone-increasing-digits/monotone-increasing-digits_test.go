package main

import "testing"

func Test_monotoneIncreasingDigits(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{10}, 9},
		{"2", args{1234}, 1234},
		{"3", args{332}, 299},
		{"4", args{4321}, 3999},
		{"5", args{2345345}, 2344999},
		{"6", args{10034}, 9999},
		{"7", args{199935}, 189999},
		{"8", args{999935}, 899999},
		{"7", args{0}, 0},
		{"8", args{789123}, 788999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := monotoneIncreasingDigits(tt.args.N); got != tt.want {
				t.Errorf("monotoneIncreasingDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

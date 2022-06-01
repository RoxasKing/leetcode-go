package main

import "testing"

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{10, 3}, 3},
		{"2", args{7, -3}, -2},
		{"3", args{-2147483648, -2147483648}, 1},
		{"4", args{0, -2147483648}, 0},
		{"5", args{-2147483648, -1}, 2147483647},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

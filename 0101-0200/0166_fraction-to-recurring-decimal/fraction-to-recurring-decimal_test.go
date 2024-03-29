package main

import "testing"

func Test_fractionToDecimal(t *testing.T) {
	type args struct {
		numerator   int
		denominator int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{1, 2}, "0.5"},
		{"2", args{2, 1}, "2"},
		{"3", args{2, 3}, "0.(6)"},
		{"4", args{-22, -2}, "11"},
		{"5", args{0, -5}, "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fractionToDecimal(tt.args.numerator, tt.args.denominator); got != tt.want {
				t.Errorf("fractionToDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

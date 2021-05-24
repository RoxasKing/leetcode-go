package main

import "testing"

func Test_strangePrinter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"aaabbb"}, 2},
		{"2", args{"aba"}, 2},
		{"3", args{"abacdc"}, 4},
		{"4", args{"acbca"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strangePrinter(tt.args.s); got != tt.want {
				t.Errorf("strangePrinter() = %v, want %v", got, tt.want)
			}
		})
	}
}

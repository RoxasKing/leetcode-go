package main

import "testing"

func Test_solveEquation(t *testing.T) {
	type args struct {
		equation string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"x+5-3+x=6+x-2"}, "x=2"},
		{"2", args{"x=x"}, "Infinite solutions"},
		{"3", args{"2x=x"}, "x=0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveEquation(tt.args.equation); got != tt.want {
				t.Errorf("solveEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}

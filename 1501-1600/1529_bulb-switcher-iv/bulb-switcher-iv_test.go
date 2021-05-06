package main

import "testing"

func Test_minFlips(t *testing.T) {
	type args struct {
		target string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"10111"}, 3},
		{"2", args{"101"}, 3},
		{"3", args{"00000"}, 0},
		{"4", args{"001011101"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFlips(tt.args.target); got != tt.want {
				t.Errorf("minFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}

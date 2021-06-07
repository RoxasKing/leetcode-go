package main

import "testing"

func Test_minFlips(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"111000"}, 2},
		{"2", args{"01001001101"}, 2},
		{"3", args{"10001100101000000"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFlips(tt.args.s); got != tt.want {
				t.Errorf("minFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}

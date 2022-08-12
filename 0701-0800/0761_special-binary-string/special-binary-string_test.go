package main

import "testing"

func Test_makeLargestSpecial(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"11011000"}, "11100100"},
		{"2", args{"10"}, "10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeLargestSpecial(tt.args.s); got != tt.want {
				t.Errorf("makeLargestSpecial() = %v, want %v", got, tt.want)
			}
		})
	}
}

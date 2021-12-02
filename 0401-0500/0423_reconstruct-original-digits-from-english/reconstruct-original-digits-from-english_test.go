package main

import "testing"

func Test_originalDigits(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"owoztneoer"}, "012"},
		{"2", args{"fviefuro"}, "45"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := originalDigits(tt.args.s); got != tt.want {
				t.Errorf("originalDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

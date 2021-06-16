package main

import "testing"

func Test_getHint(t *testing.T) {
	type args struct {
		secret string
		guess  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"1807", "7810"}, "1A3B"},
		{"2", args{"1123", "0111"}, "1A1B"},
		{"3", args{"1", "0"}, "0A0B"},
		{"4", args{"1", "1"}, "1A0B"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHint(tt.args.secret, tt.args.guess); got != tt.want {
				t.Errorf("getHint() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_splitString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"1234"}, false},
		{"2", args{"050043"}, true},
		{"3", args{"9080701"}, false},
		{"4", args{"10009998"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitString(tt.args.s); got != tt.want {
				t.Errorf("splitString() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_checkValidString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"()"}, true},
		{"2", args{"(*)"}, true},
		{"3", args{"(*))"}, true},
		{"4", args{"("}, false},
		{"5", args{")"}, false},
		{"6", args{"*"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValidString(tt.args.s); got != tt.want {
				t.Errorf("checkValidString() = %v, want %v", got, tt.want)
			}
		})
	}
}

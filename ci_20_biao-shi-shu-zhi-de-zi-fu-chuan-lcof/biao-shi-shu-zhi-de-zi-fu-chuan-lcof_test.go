package main

import "testing"

func Test_isNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"+100"}, true},
		{"2", args{"5e2"}, true},
		{"3", args{"-123"}, true},
		{"4", args{"3.1416"}, true},
		{"5", args{"0123"}, true},
		{"6", args{"12e"}, false},
		{"7", args{"1a3.14"}, false},
		{"8", args{"1.2.3"}, false},
		{"9", args{"+-5"}, false},
		{"10", args{"-1E-16"}, true},
		{"11", args{"12e+5.4"}, false},
		{"12", args{"."}, false},
		{"13", args{".5"}, true},
		{"14", args{"5."}, true},
		{"15", args{" "}, false},
		{"16", args{".1."}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.s); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

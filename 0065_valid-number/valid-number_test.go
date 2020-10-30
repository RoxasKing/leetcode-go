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
		{"1", args{"0"}, true},
		{"2", args{" 0.1 "}, true},
		{"3", args{"abc"}, false},
		{"4", args{"1 a"}, false},
		{"5", args{"2e10"}, true},
		{"6", args{" -9043 "}, true},
		{"7", args{" 1e"}, false},
		{"8", args{"e3"}, false},
		{"9", args{" 6e-1"}, true},
		{"10", args{" 99e2.5"}, false},
		{"11", args{"53.5e93"}, true},
		{"12", args{" --6"}, false},
		{"13", args{"-+3"}, false},
		{"14", args{"95a54e53"}, false},
		{"15", args{" 6e+1"}, true},
		{"16", args{".1"}, true},
		{"17", args{"1."}, true},
		{"18", args{"."}, false},
		{"19", args{"46.e3"}, true},
		{"20", args{".e3"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.s); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

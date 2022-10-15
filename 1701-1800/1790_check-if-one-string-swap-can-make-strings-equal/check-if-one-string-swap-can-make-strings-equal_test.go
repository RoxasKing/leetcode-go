package main

import "testing"

func Test_areAlmostEqual(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"bank", "kanb"}, true},
		{"2", args{"attack", "defend"}, false},
		{"3", args{"kelb", "kelb"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areAlmostEqual(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("areAlmostEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

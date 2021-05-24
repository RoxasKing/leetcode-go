package main

import "testing"

func Test_checkZeroOnes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"1101"}, true},
		{"2", args{"111000"}, false},
		{"3", args{"110100010"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkZeroOnes(tt.args.s); got != tt.want {
				t.Errorf("checkZeroOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}

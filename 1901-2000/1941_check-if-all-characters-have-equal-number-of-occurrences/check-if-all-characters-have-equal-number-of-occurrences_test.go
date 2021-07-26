package main

import "testing"

func Test_areOccurrencesEqual(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"abacbc"}, true},
		{"2", args{"aaabb"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areOccurrencesEqual(tt.args.s); got != tt.want {
				t.Errorf("areOccurrencesEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

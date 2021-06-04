package main

import "testing"

func Test_minCharacters(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"aba", "caa"}, 2},
		{"2", args{"dabadd", "cda"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCharacters(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("minCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}

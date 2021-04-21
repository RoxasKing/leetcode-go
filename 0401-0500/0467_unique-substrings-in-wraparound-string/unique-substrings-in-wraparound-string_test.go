package main

import "testing"

func Test_findSubstringInWraproundString(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"a"}, 1},
		{"2", args{"cac"}, 2},
		{"3", args{"zab"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstringInWraproundString(tt.args.p); got != tt.want {
				t.Errorf("findSubstringInWraproundString() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		s       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"abba", "dog cat cat dog"}, true},
		{"2", args{"abba", "dog cat cat fish"}, false},
		{"3", args{"aaaa", "dog cat cat dog"}, false},
		{"4", args{"abba", "dog dog dog dog"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.args.pattern, tt.args.s); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

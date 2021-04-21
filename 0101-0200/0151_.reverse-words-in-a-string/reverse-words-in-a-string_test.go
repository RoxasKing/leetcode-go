package main

import (
	"testing"
)

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"the sky is blue"}, "blue is sky the"},
		{"2", args{"  hello world!  "}, "world! hello"},
		{"3", args{"a good   example"}, "example good a"},
		{"4", args{"  Bob    Loves  Alice   "}, "Alice Loves Bob"},
		{"5", args{"Alice does not even like bob"}, "bob like even not does Alice"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_reorderSpaces(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"  this   is  a sentence "}, "this   is   a   sentence"},
		{"2", args{" practice   makes   perfect"}, "practice   makes   perfect "},
		{"3", args{"hello   world"}, "hello   world"},
		{"4", args{"  walks  udp package   into  bar a"}, "walks  udp  package  into  bar  a "},
		{"5", args{"a"}, "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderSpaces(tt.args.text); got != tt.want {
				t.Errorf("reorderSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

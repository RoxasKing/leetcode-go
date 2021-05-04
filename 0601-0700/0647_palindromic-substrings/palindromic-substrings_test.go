package main

import (
	"testing"
)

func Test_countSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"abc"}, 3},
		{"2", args{"aaa"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubstrings(tt.args.s); got != tt.want {
				t.Errorf("countSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countSubstrings2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"abc"}, 3},
		{"2", args{"aaa"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubstrings2(tt.args.s); got != tt.want {
				t.Errorf("countSubstrings2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countSubstrings3(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"abc"}, 3},
		{"2", args{"aaa"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubstrings3(tt.args.s); got != tt.want {
				t.Errorf("countSubstrings3() = %v, want %v", got, tt.want)
			}
		})
	}
}

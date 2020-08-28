package main

import (
	"testing"
)

func Test_minimumLengthEncoding(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]string{"time", "me", "bell"}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumLengthEncoding(tt.args.words); got != tt.want {
				t.Errorf("minimumLengthEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumLengthEncoding2(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]string{"time", "me", "bell"}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumLengthEncoding2(tt.args.words); got != tt.want {
				t.Errorf("minimumLengthEncoding2() = %v, want %v", got, tt.want)
			}
		})
	}
}

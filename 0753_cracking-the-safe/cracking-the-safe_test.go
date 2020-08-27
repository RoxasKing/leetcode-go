package main

import (
	"testing"
)

func Test_crackSafe(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{1, 2}, "01"},
		{"2", args{2, 2}, "00110"},
		{"3", args{2, 3}, "0221120100"},
		{"4", args{3, 2}, "0011101000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := crackSafe(tt.args.n, tt.args.k)
			if got != tt.want && got != reverse(tt.want) {
				t.Errorf("crackSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func reverse(str string) string {
	bytes := []byte(str)
	for i := 0; i < len(bytes)>>1; i++ {
		bytes[i], bytes[len(bytes)-1-i] = bytes[len(bytes)-1-i], bytes[i]
	}
	return string(bytes)
}

func Test_crackSafe2(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{1, 2}, "01"},
		{"2", args{2, 2}, "00110"},
		{"3", args{2, 3}, "0221120100"},
		{"4", args{3, 2}, "0011101000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := crackSafe(tt.args.n, tt.args.k)
			if got != tt.want && got != reverse(tt.want) {
				t.Errorf("crackSafe2() = %v, want %v", got, tt.want)
			}
		})
	}
}

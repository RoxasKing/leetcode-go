package main

import (
	"testing"
)

func Test_divisorGame(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{2}, true},
		{"2", args{3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisorGame(tt.args.N); got != tt.want {
				t.Errorf("divisorGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_divisorGame2(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{2}, true},
		{"2", args{3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisorGame2(tt.args.N); got != tt.want {
				t.Errorf("divisorGame2() = %v, want %v", got, tt.want)
			}
		})
	}
}

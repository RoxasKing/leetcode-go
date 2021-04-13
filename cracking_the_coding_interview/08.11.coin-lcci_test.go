package main

import (
	"testing"
)

func Test_waysToChange(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5}, 2},
		{"2", args{10}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToChange(tt.args.n); got != tt.want {
				t.Errorf("waysToChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_waysToChange2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5}, 2},
		{"2", args{10}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToChange2(tt.args.n); got != tt.want {
				t.Errorf("waysToChange2() = %v, want %v", got, tt.want)
			}
		})
	}
}

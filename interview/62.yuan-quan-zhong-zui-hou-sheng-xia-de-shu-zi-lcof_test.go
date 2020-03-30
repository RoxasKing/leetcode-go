package interview

import (
	"testing"
)

func Test_lastRemaining(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5, 3}, 3},
		{"2", args{10, 17}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastRemaining(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("lastRemaining() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lastRemaining2(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5, 3}, 3},
		{"2", args{10, 17}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastRemaining2(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("lastRemaining2() = %v, want %v", got, tt.want)
			}
		})
	}
}

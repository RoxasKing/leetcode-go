package main

import "testing"

func Test_rotatedDigits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{10}, 4},
		{"2", args{1}, 0},
		{"3", args{2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotatedDigits(tt.args.n); got != tt.want {
				t.Errorf("rotatedDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

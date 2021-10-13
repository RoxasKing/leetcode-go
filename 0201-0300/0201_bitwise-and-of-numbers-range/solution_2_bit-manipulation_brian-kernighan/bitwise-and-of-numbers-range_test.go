package main

import "testing"

func Test_rangeBitwiseAnd(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5, 7}, 4},
		{"2", args{0, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeBitwiseAnd2(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("rangeBitwiseAnd2() = %v, want %v", got, tt.want)
			}
		})
	}
}

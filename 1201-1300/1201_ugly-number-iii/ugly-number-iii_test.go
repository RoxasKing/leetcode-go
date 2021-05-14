package main

import "testing"

func Test_nthUglyNumber(t *testing.T) {
	type args struct {
		n int
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{3, 2, 3, 5}, 4},
		{"2", args{4, 2, 3, 4}, 6},
		{"3", args{5, 2, 11, 13}, 10},
		{"4", args{1000000000, 2, 217983653, 336916467}, 1999999984},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthUglyNumber(tt.args.n, tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("nthUglyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

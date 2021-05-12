package main

import "testing"

func Test_nthMagicalNumber(t *testing.T) {
	type args struct {
		n int
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1, 2, 3}, 2},
		{"2", args{4, 2, 3}, 6},
		{"3", args{5, 2, 4}, 10},
		{"4", args{3, 6, 4}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthMagicalNumber(tt.args.n, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("nthMagicalNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

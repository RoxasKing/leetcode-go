package main

import (
	"math"
	"testing"
)

func Test_new21Game(t *testing.T) {
	type args struct {
		N int
		K int
		W int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1", args{10, 1, 10}, 1},
		{"2", args{6, 1, 10}, 0.6},
		{"3", args{21, 17, 10}, 0.73278},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := new21Game(tt.args.N, tt.args.K, tt.args.W); math.Abs(got-tt.want) > 1e-5 {
				t.Errorf("new21Game() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"testing"
)

func Test_superEggDrop(t *testing.T) {
	type args struct {
		K int
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1, 2}, 2},
		{"2", args{2, 6}, 3},
		{"3", args{3, 14}, 4},
		{"4", args{1, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superEggDrop(tt.args.K, tt.args.N); got != tt.want {
				t.Errorf("superEggDrop() = %v, want %v", got, tt.want)
			}
		})
	}
}

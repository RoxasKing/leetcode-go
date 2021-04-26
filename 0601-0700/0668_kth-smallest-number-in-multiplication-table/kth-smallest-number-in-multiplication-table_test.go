package main

import "testing"

func Test_findKthNumber(t *testing.T) {
	type args struct {
		m int
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{3, 3, 5}, 3},
		{"2", args{2, 3, 6}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthNumber(tt.args.m, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("findKthNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

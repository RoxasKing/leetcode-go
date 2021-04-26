package main

import "testing"

func Test_sumBase(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{34, 6}, 9},
		{"2", args{10, 10}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumBase(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("sumBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

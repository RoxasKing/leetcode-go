package main

import "testing"

func Test_rearrangeSticks(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{3, 2}, 3},
		{"2", args{5, 5}, 1},
		{"3", args{20, 11}, 647427950},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rearrangeSticks(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("rearrangeSticks() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_magicalString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{6}, 3},
		{"2", args{1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := magicalString(tt.args.n); got != tt.want {
				t.Errorf("magicalString() = %v, want %v", got, tt.want)
			}
		})
	}
}

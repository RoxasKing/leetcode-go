package main

import "testing"

func Test_getKthMagicNumber(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKthMagicNumber(tt.args.k); got != tt.want {
				t.Errorf("getKthMagicNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

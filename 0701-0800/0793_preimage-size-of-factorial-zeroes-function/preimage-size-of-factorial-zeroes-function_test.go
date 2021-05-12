package main

import "testing"

func Test_preimageSizeFZF(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{0}, 5},
		{"2", args{5}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preimageSizeFZF(tt.args.k); got != tt.want {
				t.Errorf("preimageSizeFZF() = %v, want %v", got, tt.want)
			}
		})
	}
}

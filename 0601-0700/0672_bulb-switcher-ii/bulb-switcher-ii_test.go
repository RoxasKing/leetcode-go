package main

import "testing"

func Test_flipLights(t *testing.T) {
	type args struct {
		n       int
		presses int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1, 1}, 2},
		{"2", args{2, 1}, 3},
		{"3", args{3, 1}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipLights(tt.args.n, tt.args.presses); got != tt.want {
				t.Errorf("flipLights() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_mirrorReflection(t *testing.T) {
	type args struct {
		p int
		q int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{2, 1}, 2},
		{"2", args{3, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mirrorReflection(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("mirrorReflection() = %v, want %v", got, tt.want)
			}
		})
	}
}

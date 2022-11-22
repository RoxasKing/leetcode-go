package main

import "testing"

func Test_soupServings(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1", args{50}, 0.625},
		{"2", args{100}, 0.71875},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := soupServings(tt.args.n); got != tt.want {
				t.Errorf("soupServings() = %v, want %v", got, tt.want)
			}
		})
	}
}

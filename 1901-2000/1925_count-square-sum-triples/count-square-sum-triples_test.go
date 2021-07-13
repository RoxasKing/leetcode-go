package main

import "testing"

func Test_countTriples(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5}, 2},
		{"2", args{10}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTriples(tt.args.n); got != tt.want {
				t.Errorf("countTriples() = %v, want %v", got, tt.want)
			}
		})
	}
}

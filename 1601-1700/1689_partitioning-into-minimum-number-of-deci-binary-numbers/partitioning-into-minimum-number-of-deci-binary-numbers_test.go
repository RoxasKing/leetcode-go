package main

import "testing"

func Test_minPartitions(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"32"}, 3},
		{"2", args{"82734"}, 8},
		{"3", args{"27346209830709182346"}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPartitions(tt.args.n); got != tt.want {
				t.Errorf("minPartitions() = %v, want %v", got, tt.want)
			}
		})
	}
}

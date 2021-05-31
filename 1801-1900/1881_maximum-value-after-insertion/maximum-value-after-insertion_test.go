package main

import "testing"

func Test_maxValue(t *testing.T) {
	type args struct {
		n string
		x int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"99", 9}, "999"},
		{"2", args{"-13", 2}, "-123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValue(tt.args.n, tt.args.x); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

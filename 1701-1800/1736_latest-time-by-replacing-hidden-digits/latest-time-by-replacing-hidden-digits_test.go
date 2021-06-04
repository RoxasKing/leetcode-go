package main

import "testing"

func Test_maximumTime(t *testing.T) {
	type args struct {
		time string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"2?:?0"}, "23:50"},
		{"2", args{"0?:3?"}, "09:39"},
		{"3", args{"1?:22"}, "19:22"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTime(tt.args.time); got != tt.want {
				t.Errorf("maximumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

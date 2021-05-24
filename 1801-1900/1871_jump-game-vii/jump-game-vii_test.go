package main

import "testing"

func Test_canReach(t *testing.T) {
	type args struct {
		s       string
		minJump int
		maxJump int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"011010", 2, 3}, true},
		{"2", args{"01101110", 2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canReach(tt.args.s, tt.args.minJump, tt.args.maxJump); got != tt.want {
				t.Errorf("canReach() = %v, want %v", got, tt.want)
			}
		})
	}
}

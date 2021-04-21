package main

import "testing"

func Test_truncateSentence(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"Hello how are you Contestant", 4}, "Hello how are you"},
		{"2", args{"What is the solution to this problem", 4}, "What is the solution"},
		{"3", args{"chopper is not a tanuki", 5}, "chopper is not a tanuki"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := truncateSentence(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("truncateSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}

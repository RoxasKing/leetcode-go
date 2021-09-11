package main

import "testing"

func Test_orderlyQueue(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"cba", 1}, "acb"},
		{"2", args{"baaca", 3}, "aaabc"},
		{"3", args{"nhtq", 1}, "htqn"},
		{"4", args{"kuh", 1}, "hku"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orderlyQueue(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("orderlyQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

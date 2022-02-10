package main

import "testing"

func Test_longestDiverseString(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{1, 1, 7}, "ccaccbcc"},
		{"2", args{7, 1, 0}, "aabaa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestDiverseString(tt.args.a, tt.args.b, tt.args.c); len(got) != len(tt.want) {
				t.Errorf("longestDiverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

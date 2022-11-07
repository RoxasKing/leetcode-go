package main

import "testing"

func Test_maxRepeating(t *testing.T) {
	type args struct {
		sequence string
		word     string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"ababc", "ab"}, 2},
		{"2", args{"ababc", "ba"}, 1},
		{"3", args{"ababc", "ac"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRepeating(tt.args.sequence, tt.args.word); got != tt.want {
				t.Errorf("maxRepeating() = %v, want %v", got, tt.want)
			}
		})
	}
}

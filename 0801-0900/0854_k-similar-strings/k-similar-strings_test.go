package main

import "testing"

func Test_kSimilarity(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"ab", "ba"}, 1},
		{"2", args{"abc", "bca"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kSimilarity(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("kSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}

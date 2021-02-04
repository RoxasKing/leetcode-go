package main

import "testing"

func Test_numSimilarGroups(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]string{"tars", "rats", "arts", "star"}}, 2},
		{"2", args{[]string{"omv", "ovm"}}, 1},
		{"3", args{[]string{"abc", "abc"}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSimilarGroups(tt.args.strs); got != tt.want {
				t.Errorf("numSimilarGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

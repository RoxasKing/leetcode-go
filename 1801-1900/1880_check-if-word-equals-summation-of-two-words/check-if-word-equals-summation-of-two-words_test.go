package main

import "testing"

func Test_isSumEqual(t *testing.T) {
	type args struct {
		firstWord  string
		secondWord string
		targetWord string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"acb", "cba", "cdb"}, true},
		{"2", args{"aaa", "a", "aab"}, false},
		{"3", args{"aaa", "a", "aaaa"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSumEqual(tt.args.firstWord, tt.args.secondWord, tt.args.targetWord); got != tt.want {
				t.Errorf("isSumEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

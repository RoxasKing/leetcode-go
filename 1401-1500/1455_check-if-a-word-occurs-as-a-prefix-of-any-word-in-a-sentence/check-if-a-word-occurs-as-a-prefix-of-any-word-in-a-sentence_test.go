package main

import "testing"

func Test_isPrefixOfWord(t *testing.T) {
	type args struct {
		sentence   string
		searchWord string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"i love eating burger", "burg"}, 4},
		{"2", args{"this problem is an easy problem", "pro"}, 2},
		{"3", args{"i am tired", "you"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrefixOfWord(tt.args.sentence, tt.args.searchWord); got != tt.want {
				t.Errorf("isPrefixOfWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_countValidWords(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"cat and  dog"}, 3},
		{"2", args{"!this  1-s b8d!"}, 0},
		{"3", args{"alice and  bob are playing stone-game10"}, 5},
		{"4", args{"a"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countValidWords(tt.args.sentence); got != tt.want {
				t.Errorf("countValidWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

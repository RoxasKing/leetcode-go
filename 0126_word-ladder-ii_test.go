package leetcode

import (
	"reflect"
	"testing"
)

func Test_findLadders(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"",
			args{
				"red",
				"tax",
				[]string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"},
			},
			[][]string{
				{"red", "ted", "tex", "tax"},
				{"red", "ted", "tad", "tax"},
				{"red", "rex", "tex", "tax"},
			},
		},
		{
			"",
			args{
				"hit",
				"cog",
				[]string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			[][]string{
				{"hit", "hot", "dot", "dog", "cog"},
				{"hit", "hot", "lot", "log", "cog"},
			},
		},
		{
			"",
			args{
				"hit",
				"cog",
				[]string{"hot", "dot", "dog", "lot", "log"},
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLadders(tt.args.beginWord, tt.args.endWord, tt.args.wordList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findLadders() = %v, want %v", got, tt.want)
			}
		})
	}
}

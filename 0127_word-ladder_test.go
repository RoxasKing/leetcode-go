package main

import (
	"testing"
)

func Test_ladderLength(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{
				"hit",
				"cog",
				[]string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			5,
		},
		{
			"",
			args{
				"hit",
				"cog",
				[]string{"hot", "dot", "dog", "lot", "log"},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ladderLength(tt.args.beginWord, tt.args.endWord, tt.args.wordList); got != tt.want {
				t.Errorf("ladderLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ladderLength2(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{
				"hit",
				"cog",
				[]string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			5,
		},
		{
			"",
			args{
				"hit",
				"cog",
				[]string{"hot", "dot", "dog", "lot", "log"},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ladderLength2(tt.args.beginWord, tt.args.endWord, tt.args.wordList); got != tt.want {
				t.Errorf("ladderLength2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ladderLength3(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{
				"hit",
				"cog",
				[]string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			5,
		},
		{
			"",
			args{
				"hit",
				"cog",
				[]string{"hot", "dot", "dog", "lot", "log"},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ladderLength3(tt.args.beginWord, tt.args.endWord, tt.args.wordList); got != tt.want {
				t.Errorf("ladderLength3() = %v, want %v", got, tt.want)
			}
		})
	}
}

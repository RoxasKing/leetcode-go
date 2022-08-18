package main

import (
	"reflect"
	"sort"
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
			"1",
			args{
				"red",
				"tax",
				[]string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"},
			},
			[][]string{
				{"red", "rex", "tex", "tax"},
				{"red", "ted", "tex", "tax"},
				{"red", "ted", "tad", "tax"},
			},
		},
		{
			"2",
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
			"3",
			args{
				"hit",
				"cog",
				[]string{"hot", "dot", "dog", "lot", "log"},
			},
			nil,
		},
		{
			"4",
			args{
				"kiss",
				"tusk",
				[]string{"miss", "dusk", "kiss", "musk", "tusk", "diss", "disk", "sang", "ties", "muss"},
			},
			[][]string{{"kiss", "diss", "disk", "dusk", "tusk"}, {"kiss", "miss", "muss", "musk", "tusk"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findLadders(tt.args.beginWord, tt.args.endWord, tt.args.wordList)
			sort.Slice(got, func(i, j int) bool {
				a, b := got[i], got[j]
				for k := range got {
					if a[k] != b[k] {
						return a[k] < b[k]
					}
				}
				return false
			})
			sort.Slice(tt.want, func(i, j int) bool {
				a, b := tt.want[i], tt.want[j]
				for k := range got {
					if a[k] != b[k] {
						return a[k] < b[k]
					}
				}
				return false
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findLadders() = %v, want %v", got, tt.want)
			}
		})
	}
}

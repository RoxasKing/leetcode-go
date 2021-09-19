package main

import "testing"

func Test_findWords(t *testing.T) {
	type args struct {
		board [][]byte
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"1",
			args{
				[][]byte{
					{'o', 'a', 'a', 'n'},
					{'e', 't', 'a', 'e'},
					{'i', 'h', 'k', 'r'},
					{'i', 'f', 'l', 'v'},
				},
				[]string{"oath", "pea", "eat", "rain"},
			},
			[]string{"eat", "oath"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findWords(tt.args.board, tt.args.words)
			src := make(map[string]bool)
			dst := make(map[string]bool)
			for _, w := range got {
				src[w] = true
			}
			for _, w := range tt.want {
				dst[w] = true
			}
			for _, w := range got {
				if !dst[w] {
					t.Errorf("findWords() = %v, want %v", got, tt.want)
				}
			}
			for _, w := range tt.want {
				if !src[w] {
					t.Errorf("findWords() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

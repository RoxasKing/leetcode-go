package main

import "testing"

func Test_isAlienSorted(t *testing.T) {
	type args struct {
		words []string
		order string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"}, true},
		{"2", args{[]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz"}, false},
		{"3", args{[]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"}, false},
		{"4", args{[]string{"kuvp", "q"}, "ngxlkthsjuoqcpavbfdermiywz"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAlienSorted(tt.args.words, tt.args.order); got != tt.want {
				t.Errorf("isAlienSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}

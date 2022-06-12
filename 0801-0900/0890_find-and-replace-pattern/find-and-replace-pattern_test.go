package main

import (
	"reflect"
	"testing"
)

func Test_findAndReplacePattern(t *testing.T) {
	type args struct {
		words   []string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{[]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb"}, []string{"mee", "aqq"}},
		{"2", args{[]string{"a", "b", "c"}, "a"}, []string{"a", "b", "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAndReplacePattern(tt.args.words, tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAndReplacePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

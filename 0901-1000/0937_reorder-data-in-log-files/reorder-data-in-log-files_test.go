package main

import (
	"reflect"
	"testing"
)

func Test_reorderLogFiles(t *testing.T) {
	type args struct {
		logs []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{[]string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}}, []string{"let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"}},
		{"2", args{[]string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"}}, []string{"g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"}},
		{"3", args{[]string{"1 n u", "r 527", "j 893", "6 14", "6 82"}}, []string{"1 n u", "r 527", "j 893", "6 14", "6 82"}},
		{"4", args{[]string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo", "a2 act car"}}, []string{"a2 act car", "g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"}},
		{"5", args{[]string{"dig1 8 1 5 1", "let1 art zero can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}}, []string{"let3 art zero", "let1 art zero can", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderLogFiles(tt.args.logs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reorderLogFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

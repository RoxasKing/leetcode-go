package main

import (
	"reflect"
	"testing"
)

func Test_stringMatching(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{[]string{"mass", "as", "hero", "superhero"}}, []string{"hero", "as"}},
		{"2", args{[]string{"leetcode", "et", "code"}}, []string{"code", "et"}},
		{"3", args{[]string{"blue", "green", "bu"}}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringMatching(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringMatching() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_uncommonFromSentences(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{"this apple is sweet", "this apple is sour"}, []string{"sweet", "sour"}},
		{"2", args{"apple apple", "banana"}, []string{"banana"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uncommonFromSentences(tt.args.s1, tt.args.s2)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uncommonFromSentences() = %v, want %v", got, tt.want)
			}
		})
	}
}

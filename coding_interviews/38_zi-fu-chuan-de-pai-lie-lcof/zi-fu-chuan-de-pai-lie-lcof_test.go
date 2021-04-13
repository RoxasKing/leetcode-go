package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_permutation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"1",
			args{"abc"},
			[]string{"abc", "acb", "bac", "bca", "cba", "cab"},
		},
		{
			"2",
			args{"aab"},
			[]string{"aab", "aba", "baa"},
		},
		{
			"3",
			args{"aaab"},
			[]string{"aaab", "aaba", "abaa", "baaa"},
		},
		{
			"4",
			args{"bccc"},
			[]string{"bccc", "cbcc", "ccbc", "cccb"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permutation(tt.args.s)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

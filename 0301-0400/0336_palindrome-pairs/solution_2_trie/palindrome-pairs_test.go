package main

import (
	"reflect"
	"testing"
)

func Test_palindromePairs(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{[]string{"abcd", "dcba", "lls", "s", "sssll"}},
			[][]int{{0, 1}, {1, 0}, {3, 2}, {2, 4}},
		},
		{
			"2",
			args{[]string{"bat", "tab", "cat"}},
			[][]int{{0, 1}, {1, 0}},
		},
		{"3", args{[]string{"a", ""}}, [][]int{{0, 1}, {1, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindromePairs(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("palindromePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

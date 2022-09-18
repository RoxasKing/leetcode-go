package main

import (
	"reflect"
	"sort"
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
			got := palindromePairs(tt.args.words)
			sort.Slice(got, func(i, j int) bool {
				return got[i][0] < got[j][0] || got[i][0] == got[j][0] && got[i][1] < got[j][1]
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return tt.want[i][0] < tt.want[j][0] || tt.want[i][0] == tt.want[j][0] && tt.want[i][1] < tt.want[j][1]
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("palindromePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

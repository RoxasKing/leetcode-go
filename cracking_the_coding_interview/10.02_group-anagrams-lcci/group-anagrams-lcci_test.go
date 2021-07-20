package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"1", args{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			[][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(sortStrs(got), sortStrs(tt.want)) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sortStrs(strs [][]string) [][]string {
	for i := range strs {
		sort.Strings(strs[i])
	}
	sort.Slice(strs, func(i, j int) bool {
		if len(strs[i]) != len(strs[j]) {
			return len(strs[i]) < len(strs[j])
		}
		return strs[i][0] < strs[j][0]
	})
	return strs
}

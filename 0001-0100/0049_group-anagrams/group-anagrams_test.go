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
			"1",
			args{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			[][]string{
				{"ate", "eat", "tea"},
				{"nat", "tan"},
				{"bat"},
			},
		},
		{"2", args{[]string{""}}, [][]string{{""}}},
		{"3", args{[]string{"a"}}, [][]string{{"a"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !compare(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compare(a, b [][]string) bool {
	sort.Slice(a, func(i, j int) bool {
		return len(a[i]) > len(a[j])
	})
	for i := range a {
		sort.Strings(a[i])
	}
	sort.Slice(b, func(i, j int) bool {
		return len(b[i]) > len(b[j])
	})
	for i := range b {
		sort.Strings(b[i])
	}
	return reflect.DeepEqual(a, b)
}

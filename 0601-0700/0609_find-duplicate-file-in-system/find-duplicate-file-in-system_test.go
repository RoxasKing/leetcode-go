package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findDuplicate(t *testing.T) {
	type args struct {
		paths []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"1", args{[]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"}}, [][]string{{"root/a/2.txt", "root/c/d/4.txt", "root/4.txt"}, {"root/a/1.txt", "root/c/3.txt"}}},
		{"2", args{[]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)"}}, [][]string{{"root/a/2.txt", "root/c/d/4.txt"}, {"root/a/1.txt", "root/c/3.txt"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args.paths); !reflect.DeepEqual(sortStrs(got), sortStrs(tt.want)) {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
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

package main

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicateFolder(t *testing.T) {
	type args struct {
		paths [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"1",
			args{[][]string{{"a"}, {"c"}, {"d"}, {"a", "b"}, {"c", "b"}, {"d", "a"}}},
			[][]string{{"d"}, {"d", "a"}},
		},
		{
			"2",
			args{[][]string{{"a"}, {"c"}, {"a", "b"}, {"c", "b"}, {"a", "b", "x"}, {"a", "b", "x", "y"}, {"w"}, {"w", "y"}}},
			[][]string{{"c"}, {"c", "b"}, {"a"}, {"a", "b"}},
		},
		{
			"3",
			args{[][]string{{"a", "b"}, {"c", "d"}, {"c"}, {"a"}}},
			[][]string{{"c"}, {"c", "d"}, {"a"}, {"a", "b"}},
		},
		{
			"4",
			args{[][]string{{"a"}, {"a", "x"}, {"a", "x", "y"}, {"a", "z"}, {"b"}, {"b", "x"}, {"b", "x", "y"}, {"b", "z"}}},
			[][]string{},
		},
		{
			"5",
			args{[][]string{{"a"}, {"a", "x"}, {"a", "x", "y"}, {"a", "z"}, {"b"}, {"b", "x"}, {"b", "x", "y"}, {"b", "z"}, {"b", "w"}}},
			[][]string{{"b"}, {"b", "w"}, {"b", "z"}, {"a"}, {"a", "z"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicateFolder(tt.args.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicateFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}

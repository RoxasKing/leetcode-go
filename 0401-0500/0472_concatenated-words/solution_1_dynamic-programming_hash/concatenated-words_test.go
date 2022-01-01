package main

import (
	"reflect"
	"testing"
)

func Test_findAllConcatenatedWordsInADict(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"1",
			args{[]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}},
			[]string{"catsdogcats", "dogcatsdog", "ratcatdogcat"},
		},
		{
			"2",
			args{[]string{"cat", "dog", "catdog"}},
			[]string{"catdog"},
		},
		{"3", args{[]string{""}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAllConcatenatedWordsInADict(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllConcatenatedWordsInADict() = %v, want %v", got, tt.want)
			}
		})
	}
}

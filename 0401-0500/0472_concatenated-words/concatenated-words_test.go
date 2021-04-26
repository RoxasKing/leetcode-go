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
		{
			"3",
			args{[]string{""}},
			[]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAllConcatenatedWordsInADict(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllConcatenatedWordsInADict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findAllConcatenatedWordsInADict2(t *testing.T) {
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
		{
			"3",
			args{[]string{""}},
			[]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAllConcatenatedWordsInADict2(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllConcatenatedWordsInADict2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_findWords(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{[]string{"Hello", "Alaska", "Dad", "Peace"}}, []string{"Alaska", "Dad"}},
		{"2", args{[]string{"omk"}}, []string{}},
		{"3", args{[]string{"adsdf", "sfd"}}, []string{"adsdf", "sfd"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWords(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

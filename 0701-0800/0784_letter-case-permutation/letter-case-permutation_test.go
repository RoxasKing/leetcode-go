package main

import (
	"reflect"
	"testing"
)

func Test_letterCasePermutation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{"a1b2"}, []string{"a1b2", "a1B2", "A1b2", "A1B2"}},
		{"2", args{"3z4"}, []string{"3z4", "3Z4"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCasePermutation(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCasePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

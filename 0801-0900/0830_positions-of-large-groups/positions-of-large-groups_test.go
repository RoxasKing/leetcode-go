package main

import (
	"reflect"
	"testing"
)

func Test_largeGroupPositions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{"abbxxxxzzy"}, [][]int{{3, 6}}},
		{"2", args{"abc"}, [][]int{}},
		{"3", args{"abcdddeeeeaabbbcd"}, [][]int{{3, 5}, {6, 9}, {12, 14}}},
		{"4", args{"aba"}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largeGroupPositions(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largeGroupPositions() = %v, want %v", got, tt.want)
			}
		})
	}
}

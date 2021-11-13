package main

import (
	"reflect"
	"testing"
)

func Test_findNumOfValidWords(t *testing.T) {
	type args struct {
		words   []string
		puzzles []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1",
			args{
				[]string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"},
				[]string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"},
			},
			[]int{1, 1, 3, 2, 4, 0},
		},
		{
			"2",
			args{
				[]string{"apple", "pleas", "please"},
				[]string{"aelwxyz", "aelpxyz", "aelpsxy", "saelpxy", "xaelpsy"},
			},
			[]int{0, 1, 3, 2, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumOfValidWords(tt.args.words, tt.args.puzzles); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumOfValidWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

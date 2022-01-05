package main

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"1",
			args{"efe"},
			[][]string{
				{"e", "f", "e"},
				{"efe"},
			},
		},
		{
			"2",
			args{"aab"},
			[][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

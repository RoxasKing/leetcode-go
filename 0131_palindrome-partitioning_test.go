package leetcode

import (
	"reflect"
	"testing"
)

func Test_partition0131(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"",
			args{"efe"},
			[][]string{
				{"e", "f", "e"},
				{"efe"},
			},
		},
		{
			"",
			args{"aab"},
			[][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition0131(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition0131() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partition01312(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"",
			args{"efe"},
			[][]string{
				{"e", "f", "e"},
				{"efe"},
			},
		},
		{
			"",
			args{"aab"},
			[][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition01312(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition01312() = %v, want %v", got, tt.want)
			}
		})
	}
}

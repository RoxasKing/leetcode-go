package My_LeetCode_In_Go

import (
	"reflect"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"",
			args{
				"wordgoodgoodgoodbestword",
				[]string{"word", "good", "best", "word"},
			},
			nil,
		},
		{
			"",
			args{
				"wordgoodgoodgoodbestword",
				[]string{"word", "good", "best", "good"},
			},
			[]int{8},
		},
		{
			"",
			args{
				"barfoothefoobarman",
				[]string{"foo", "bar"},
			},
			[]int{0, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstring(tt.args.s, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

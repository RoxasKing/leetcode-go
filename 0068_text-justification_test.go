package My_LeetCode_In_Go

import (
	"reflect"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"",
			args{
				[]string{
					"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country",
				},
				16,
			},
			[]string{
				"ask   not   what",
				"your country can",
				"do  for  you ask",
				"what  you can do",
				"for your country",
			},
		},
		{
			"",
			args{
				[]string{
					"This", "is", "an", "example", "of", "text", "justification.",
				},
				16,
			},
			[]string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		},
		{
			"",
			args{
				[]string{
					"What", "must", "be", "acknowledgment", "shall", "be",
				},
				16,
			},
			[]string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		},
		{
			"",
			args{
				[]string{
					"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do",
				},
				20,
			},
			[]string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify2(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("fullJustify() = %v, want %v", got, tt.want)
			}
		})
	}
}

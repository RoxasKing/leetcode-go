package main

import (
	"reflect"
	"testing"
)

func Test_removeComments(t *testing.T) {
	type args struct {
		source []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"1",
			args{
				[]string{
					"/*Test program */",
					"int main()",
					"{ ",
					"  // variable declaration ",
					"int a, b, c;",
					"/* This is a test",
					"   multiline  ",
					"   comment for ",
					"   testing */",
					"a = b + c;", "}",
				},
			},
			[]string{"int main()", "{ ", "  ", "int a, b, c;", "a = b + c;", "}"},
		},
		{
			"2",
			args{
				[]string{"a/*comment", "line", "more_comment*/b"},
			},
			[]string{"ab"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeComments(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

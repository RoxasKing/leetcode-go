package main

import "testing"

func Test_evaluate(t *testing.T) {
	type args struct {
		s         string
		knowledge [][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				"(name)is(age)yearsold",
				[][]string{{"name", "bob"}, {"age", "two"}},
			},
			"bobistwoyearsold",
		},
		{
			"2",
			args{
				"hi(name)",
				[][]string{{"a", "b"}},
			},
			"hi?",
		},
		{
			"3",
			args{
				"(a)(a)(a)aaa",
				[][]string{{"a", "yes"}},
			},
			"yesyesyesaaa",
		},
		{
			"4",
			args{
				"(a)(b)",
				[][]string{{"a", "b"}, {"b", "a"}},
			},
			"ba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluate(tt.args.s, tt.args.knowledge); got != tt.want {
				t.Errorf("evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}

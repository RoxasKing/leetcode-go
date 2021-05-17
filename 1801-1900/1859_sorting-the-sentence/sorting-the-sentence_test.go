package main

import "testing"

func Test_sortSentence(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"is2 sentence4 This1 a3"}, "This is a sentence"},
		{"2", args{"Myself2 Me1 I4 and3"}, "Me Myself and I"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortSentence(tt.args.s); got != tt.want {
				t.Errorf("sortSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}

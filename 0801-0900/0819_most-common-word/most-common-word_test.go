package main

import "testing"

func Test_mostCommonWord(t *testing.T) {
	type args struct {
		paragraph string
		banned    []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}}, "ball"},
		{"2", args{"a.", []string{}}, "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostCommonWord(tt.args.paragraph, tt.args.banned); got != tt.want {
				t.Errorf("mostCommonWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

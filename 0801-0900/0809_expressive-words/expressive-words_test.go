package main

import "testing"

func Test_expressiveWords(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"heeellooo", []string{"hello", "hi", "helo"}}, 1},
		{"2", args{"zzzzzyyyyy", []string{"zzyy", "zy", "zyy"}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expressiveWords(tt.args.s, tt.args.words); got != tt.want {
				t.Errorf("expressiveWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

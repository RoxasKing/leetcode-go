package main

import "testing"

func Test_countVowelStrings(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1}, 5},
		{"2", args{2}, 15},
		{"3", args{33}, 66045},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countVowelStrings(tt.args.n); got != tt.want {
				t.Errorf("countVowelStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_longestStrChain(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]string{"a", "b", "ba", "bca", "bda", "bdca"}}, 4},
		{"2", args{[]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}}, 5},
		{"3", args{[]string{"abcd", "dbqca"}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestStrChain(tt.args.words); got != tt.want {
				t.Errorf("longestStrChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

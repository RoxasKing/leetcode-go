package main

import "testing"

func Test_arrayStringsAreEqual(t *testing.T) {
	type args struct {
		word1 []string
		word2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]string{"ab", "c"}, []string{"a", "bc"}}, true},
		{"2", args{[]string{"a", "cb"}, []string{"ab", "c"}}, false},
		{"3", args{[]string{"abc", "d", "defg"}, []string{"abcddefg"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayStringsAreEqual(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("arrayStringsAreEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

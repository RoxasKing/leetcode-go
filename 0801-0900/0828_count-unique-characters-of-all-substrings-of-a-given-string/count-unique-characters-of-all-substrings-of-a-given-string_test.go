package main

import "testing"

func Test_uniqueLetterString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"ABC"}, 10},
		{"2", args{"ABA"}, 8},
		{"3", args{"LEETCODE"}, 92},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueLetterString(tt.args.s); got != tt.want {
				t.Errorf("uniqueLetterString() = %v, want %v", got, tt.want)
			}
		})
	}
}

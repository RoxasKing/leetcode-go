package main

import "testing"

func Test_distinctSubseqII(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"abc"}, 7},
		{"2", args{"aba"}, 6},
		{"3", args{"aaa"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distinctSubseqII(tt.args.S); got != tt.want {
				t.Errorf("distinctSubseqII() = %v, want %v", got, tt.want)
			}
		})
	}
}

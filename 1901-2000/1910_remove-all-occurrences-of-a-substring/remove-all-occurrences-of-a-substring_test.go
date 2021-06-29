package main

import "testing"

func Test_removeOccurrences(t *testing.T) {
	type args struct {
		s    string
		part string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"daabcbaabcbc", "abc"}, "dab"},
		{"2", args{"axxxxyyyyb", "xy"}, "ab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeOccurrences(tt.args.s, tt.args.part); got != tt.want {
				t.Errorf("removeOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_longestDupSubstring(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"banana"}, "ana"},
		{"2", args{"abcd"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestDupSubstring(tt.args.S); got != tt.want {
				t.Errorf("longestDupSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

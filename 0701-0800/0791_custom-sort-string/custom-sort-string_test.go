package main

import "testing"

func Test_customSortString(t *testing.T) {
	type args struct {
		order string
		s     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"cba", "abcd"}, "cbad"},
		{"2", args{"cbafg", "abcd"}, "cbad"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := customSortString(tt.args.order, tt.args.s); got != tt.want {
				t.Errorf("customSortString() = %v, want %v", got, tt.want)
			}
		})
	}
}

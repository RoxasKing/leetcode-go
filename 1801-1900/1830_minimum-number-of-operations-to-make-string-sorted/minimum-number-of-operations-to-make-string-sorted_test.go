package main

import "testing"

func Test_makeStringSorted(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"cba"}, 5},
		{"2", args{"aabaa"}, 2},
		{"3", args{"cdbea"}, 63},
		{"4", args{"leetcodeleetcodeleetcode"}, 982157772},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeStringSorted(tt.args.s); got != tt.want {
				t.Errorf("makeStringSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}

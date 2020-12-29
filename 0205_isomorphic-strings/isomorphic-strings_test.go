package main

import "testing"

func Test_isIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"egg", "add"}, true},
		{"2", args{"foo", "bar"}, false},
		{"3", args{"paper", "title"}, true},
		{"4", args{"ab", "aa"}, false},
		{"5", args{"ab", "ca"}, true},
		{"6", args{"13", "42"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}

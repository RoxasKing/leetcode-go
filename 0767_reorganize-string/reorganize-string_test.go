package main

import "testing"

func Test_reorganizeString(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"aab"}, "aba"},
		{"2", args{"aaab"}, ""},
		{"3", args{"aaabb"}, "ababa"},
		{"4", args{"aaabbccc"}, "acbacbac"},
		{"5", args{"vvvlo"}, "vlvov"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reorganizeString(tt.args.S)
			if len(got) != len(tt.want) {
				t.Errorf("reorganizeString() test failed")
				return
			}
			for i := 0; i < len(got)-1; i++ {
				if got[i] == got[i+1] {
					t.Errorf("reorganizeString() test failed")
					return
				}
			}
		})
	}
}

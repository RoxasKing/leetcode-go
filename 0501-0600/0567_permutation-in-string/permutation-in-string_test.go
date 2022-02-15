package main

import "testing"

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"ab", "eidbaooo"}, true},
		{"2", args{"eidbaooo", "ab"}, false},
		{"3", args{"ab", "eidboaoo"}, false},
		{"4", args{"hello", "ooolleoooleh"}, false},
		{"5", args{"ky", "ainwkckifykxlribaypk"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}

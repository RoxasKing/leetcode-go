package main

import "testing"

func Test_checkIfPangram(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"thequickbrownfoxjumpsoverthelazydog"}, true},
		{"2", args{"leetcode"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfPangram(tt.args.sentence); got != tt.want {
				t.Errorf("checkIfPangram() = %v, want %v", got, tt.want)
			}
		})
	}
}

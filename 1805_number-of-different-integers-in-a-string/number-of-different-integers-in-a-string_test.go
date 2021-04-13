package main

import "testing"

func Test_numDifferentIntegers(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"a123bc34d8ef34"}, 3},
		{"2", args{"leet1234code234"}, 2},
		{"3", args{"a1b01c001"}, 1},
		{"4", args{"0a0"}, 1},
		{"5", args{"035985750011523523129774573439111590559325"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDifferentIntegers(tt.args.word); got != tt.want {
				t.Errorf("numDifferentIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}

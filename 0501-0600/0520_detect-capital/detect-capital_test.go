package main

import "testing"

func Test_detectCapitalUse(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"USA"}, true},
		{"2", args{"FlaG"}, false},
		{"3", args{"Usa"}, true},
		{"4", args{"UsA"}, false},
		{"5", args{"U"}, true},
		{"6", args{"u"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCapitalUse(tt.args.word); got != tt.want {
				t.Errorf("detectCapitalUse() = %v, want %v", got, tt.want)
			}
		})
	}
}

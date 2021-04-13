package main

import "testing"

func Test_squareIsWhite(t *testing.T) {
	type args struct {
		coordinates string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"a1"}, false},
		{"2", args{"h3"}, true},
		{"3", args{"c7"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squareIsWhite(tt.args.coordinates); got != tt.want {
				t.Errorf("squareIsWhite() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"testing"
)

func Test_longestDecomposition(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"ghiabcdefhelloadamhelloabcdefghi"}, 7},
		{"2", args{"merchant"}, 1},
		{"3", args{"antaprezatepzapreanta"}, 11},
		{"4", args{"aaa"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestDecomposition(tt.args.text); got != tt.want {
				t.Errorf("longestDecomposition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestDecomposition2(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"ghiabcdefhelloadamhelloabcdefghi"}, 7},
		{"2", args{"merchant"}, 1},
		{"3", args{"antaprezatepzapreanta"}, 11},
		{"4", args{"aaa"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestDecomposition2(tt.args.text); got != tt.want {
				t.Errorf("longestDecomposition2() = %v, want %v", got, tt.want)
			}
		})
	}
}

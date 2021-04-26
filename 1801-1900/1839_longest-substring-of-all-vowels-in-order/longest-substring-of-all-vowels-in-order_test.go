package main

import "testing"

func Test_longestBeautifulSubstring(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"aeiaaioaaaaeiiiiouuuooaauuaeiu"}, 13},
		{"2", args{"aeeeiiiioooauuuaeiou"}, 5},
		{"3", args{"a"}, 0},
		{"4", args{"eauoiouieaaoueiuaieoeauoiaueoiaeoiuieuaoiaeouiaueo"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestBeautifulSubstring(tt.args.word); got != tt.want {
				t.Errorf("longestBeautifulSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

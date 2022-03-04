package main

import "testing"

func Test_nearestPalindromic(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"123"}, "121"},
		{"2", args{"1"}, "0"},
		{"3", args{"1234"}, "1221"},
		{"4", args{"999"}, "1001"},
		{"5", args{"1000"}, "999"},
		{"6", args{"12932"}, "12921"},
		{"7", args{"99800"}, "99799"},
		{"8", args{"12120"}, "12121"},
		{"9", args{"10"}, "9"},
		{"10", args{"11"}, "9"},
		{"11", args{"11011"}, "11111"},
		{"12", args{"1283"}, "1331"},
		{"13", args{"88"}, "77"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nearestPalindromic(tt.args.n); got != tt.want {
				t.Errorf("nearestPalindromic() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_largestOddNumber(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"52"}, "5"},
		{"2", args{"4206"}, ""},
		{"3", args{"35427"}, "35427"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestOddNumber(tt.args.num); got != tt.want {
				t.Errorf("largestOddNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

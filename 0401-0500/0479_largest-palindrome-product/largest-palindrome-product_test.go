package main

import "testing"

func Test_largestPalindrome(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1}, 9},
		{"2", args{2}, 987},
		{"3", args{3}, 123},
		{"4", args{4}, 597},
		{"5", args{5}, 677},
		{"6", args{6}, 1218},
		{"7", args{7}, 877},
		{"8", args{8}, 475},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestPalindrome(tt.args.n); got != tt.want {
				t.Errorf("largestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

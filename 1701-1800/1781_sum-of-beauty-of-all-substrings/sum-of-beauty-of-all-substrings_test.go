package main

import "testing"

func Test_beautySum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"aabcb"}, 5},
		{"2", args{"aabcbaa"}, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := beautySum(tt.args.s); got != tt.want {
				t.Errorf("beautySum() = %v, want %v", got, tt.want)
			}
		})
	}
}

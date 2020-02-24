package leetcode

import "testing"

func Test_isPalindrome0125(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"A man, a plan, a canal: Panama"}, true},
		{"", args{"race a car"}, false},
		{"", args{"0P"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome0125(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome0125() = %v, want %v", got, tt.want)
			}
		})
	}
}

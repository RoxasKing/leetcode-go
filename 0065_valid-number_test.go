package leetcode

import "testing"

func Test_isNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"0"}, true},
		{"", args{" 0.1 "}, true},
		{"", args{"abc"}, false},
		{"", args{"1 a"}, false},
		{"", args{"2e10"}, true},
		{"", args{" -9043 "}, true},
		{"", args{" 1e"}, false},
		{"", args{"e3"}, false},
		{"", args{" 6e-1"}, true},
		{"", args{" 99e2.5"}, false},
		{"", args{"53.5e93"}, true},
		{"", args{" --6"}, false},
		{"", args{"-+3"}, false},
		{"", args{"95a54e53"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.s); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

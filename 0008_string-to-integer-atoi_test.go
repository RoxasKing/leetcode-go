package leetcode

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"0-1"}, 0},
		{"", args{" -42"}, -42},
		{"", args{"+1"}, 1},
		{"", args{"-91283472332"}, -2147483648},
		{"", args{"words and 987"}, 0},
		{"", args{"4193 with words"}, 4193},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

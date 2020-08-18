package main

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
		{"1", args{"0-1"}, 0},
		{"2", args{" -42"}, -42},
		{"3", args{"+1"}, 1},
		{"4", args{"-91283472332"}, -2147483648},
		{"5", args{"words and 987"}, 0},
		{"6", args{"4193 with words"}, 4193},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

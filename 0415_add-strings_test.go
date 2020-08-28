package main

import "testing"

func Test_addStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"0", "0"}, "0"},
		{"2", args{"1", "9"}, "10"},
		{"3", args{"123", "567"}, "690"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

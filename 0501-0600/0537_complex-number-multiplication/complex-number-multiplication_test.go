package main

import "testing"

func Test_complexNumberMultiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"1+1i", "1+1i"}, "0+2i"},
		{"2", args{"1+-1i", "1+-1i"}, "0+-2i"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := complexNumberMultiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("complexNumberMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

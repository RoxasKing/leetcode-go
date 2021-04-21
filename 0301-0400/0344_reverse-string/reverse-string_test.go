package main

import "testing"

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{[]byte{'h', 'e', 'l', 'l', 'o'}}},
		{"2", args{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := make([]byte, len(tt.args.s))
			copy(before, tt.args.s)
			reverseString(tt.args.s)
			for i := 0; i < len(before); i++ {
				if before[i] != tt.args.s[len(before)-1-i] {
					t.Errorf("reverseString() test failed")
				}
			}
		})
	}
}

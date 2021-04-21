package main

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"1", args{[]byte{'h', 'e', 'l', 'l', 'o'}}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{"2", args{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			if !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("reverseString() = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}

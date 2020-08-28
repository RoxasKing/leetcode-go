package main

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"1", args{[]byte{'c', 'f', 'j'}, 'a'}, 'c'},
		{"2", args{[]byte{'c', 'f', 'j'}, 'c'}, 'f'},
		{"3", args{[]byte{'c', 'f', 'j'}, 'd'}, 'f'},
		{"4", args{[]byte{'c', 'f', 'j'}, 'g'}, 'j'},
		{"5", args{[]byte{'c', 'f', 'j'}, 'j'}, 'c'},
		{"6", args{[]byte{'c', 'f', 'j'}, 'k'}, 'c'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

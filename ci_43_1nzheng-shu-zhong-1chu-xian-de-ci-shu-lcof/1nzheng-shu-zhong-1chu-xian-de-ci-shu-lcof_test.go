package main

import (
	"testing"
)

func Test_countDigitOne(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{12}, 5},
		{"2", args{13}, 6},
		{"3", args{100}, 21},
		{"4", args{800}, 260},
		{"5", args{8000}, 3400},
		{"6", args{876}, 278},
		{"6", args{1000}, 301},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDigitOne(tt.args.n); got != tt.want {
				t.Errorf("countDigitOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

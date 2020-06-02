package codinginterviews

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDigitOne(tt.args.n); got != tt.want {
				t.Errorf("countDigitOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countDigitOne2(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDigitOne2(tt.args.n); got != tt.want {
				t.Errorf("countDigitOne2() = %v, want %v", got, tt.want)
			}
		})
	}
}

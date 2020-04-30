package leetcode

import (
	"testing"
)

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{19}, true},
		{"2", args{20}, false},
		{"3", args{21}, false},
		{"4", args{1}, true},
		{"5", args{7}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isHappy2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{19}, true},
		{"2", args{20}, false},
		{"3", args{21}, false},
		{"4", args{1}, true},
		{"5", args{7}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy2(tt.args.n); got != tt.want {
				t.Errorf("isHappy2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_isUgly(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{6}, true},
		{"2", args{8}, true},
		{"3", args{14}, false},
		{"4", args{1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUgly(tt.args.n); got != tt.want {
				t.Errorf("isUgly() = %v, want %v", got, tt.want)
			}
		})
	}
}

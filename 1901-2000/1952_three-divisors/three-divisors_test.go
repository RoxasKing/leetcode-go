package main

import "testing"

func Test_isThree(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{2}, false},
		{"2", args{4}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isThree(tt.args.n); got != tt.want {
				t.Errorf("isThree() = %v, want %v", got, tt.want)
			}
		})
	}
}

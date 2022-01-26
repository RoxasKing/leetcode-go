package main

import "testing"

func Test_numberOfMatches(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{7}, 6},
		{"2", args{14}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfMatches(tt.args.n); got != tt.want {
				t.Errorf("numberOfMatches() = %v, want %v", got, tt.want)
			}
		})
	}
}

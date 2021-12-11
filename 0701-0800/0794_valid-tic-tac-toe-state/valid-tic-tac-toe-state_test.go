package main

import "testing"

func Test_validTicTacToe(t *testing.T) {
	type args struct {
		board []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]string{"O  ", "   ", "   "}}, false},
		{"2", args{[]string{"XOX", " X ", "   "}}, false},
		{"3", args{[]string{"XXX", "   ", "OOO"}}, false},
		{"4", args{[]string{"XOX", "O O", "XOX"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validTicTacToe(tt.args.board); got != tt.want {
				t.Errorf("validTicTacToe() = %v, want %v", got, tt.want)
			}
		})
	}
}

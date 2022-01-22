package main

import "testing"

func Test_winnerSquareGame(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{1}, true},
		{"2", args{2}, false},
		{"3", args{4}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := winnerSquareGame(tt.args.n); got != tt.want {
				t.Errorf("winnerSquareGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

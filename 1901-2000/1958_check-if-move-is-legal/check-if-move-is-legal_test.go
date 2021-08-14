package main

import "testing"

func Test_checkMove(t *testing.T) {
	type args struct {
		board [][]byte
		rMove int
		cMove int
		color byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[][]byte{{'.', '.', '.', 'B', '.', '.', '.', '.'}, {'.', '.', '.', 'W', '.', '.', '.', '.'}, {'.', '.', '.', 'W', '.', '.', '.', '.'}, {'.', '.', '.', 'W', '.', '.', '.', '.'}, {'W', 'B', 'B', '.', 'W', 'W', 'W', 'B'}, {'.', '.', '.', 'B', '.', '.', '.', '.'}, {'.', '.', '.', 'B', '.', '.', '.', '.'}, {'.', '.', '.', 'W', '.', '.', '.', '.'}}, 4, 3, 'B'}, true},
		{"2", args{[][]byte{{'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', 'B', '.', '.', 'W', '.', '.', '.'}, {'.', '.', 'W', '.', '.', '.', '.', '.'}, {'.', '.', '.', 'W', 'B', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', 'B', 'W', '.', '.'}, {'.', '.', '.', '.', '.', '.', 'W', '.'}, {'.', '.', '.', '.', '.', '.', '.', 'B'}}, 4, 4, 'W'}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkMove(tt.args.board, tt.args.rMove, tt.args.cMove, tt.args.color); got != tt.want {
				t.Errorf("checkMove() = %v, want %v", got, tt.want)
			}
		})
	}
}

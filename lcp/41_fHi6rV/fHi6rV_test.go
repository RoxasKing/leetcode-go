package main

import "testing"

func Test_flipChess(t *testing.T) {
	type args struct {
		chessboard []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]string{"....X.", "....X.", "XOOO..", "......", "......"}}, 3},
		{"2", args{[]string{".X.", ".O.", "XO."}}, 2},
		{"3", args{[]string{".......", ".......", ".......", "X......", ".O.....", "..O....", "....OOX"}}, 4},
		{"4", args{[]string{".O.....", ".O.....", "XOO.OOX", ".OO.OO.", ".XO.OX.", "..X.X.."}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipChess(tt.args.chessboard); got != tt.want {
				t.Errorf("flipChess() = %v, want %v", got, tt.want)
			}
		})
	}
}

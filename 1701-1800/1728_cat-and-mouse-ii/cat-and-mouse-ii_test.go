package main

import "testing"

func Test_canMouseWin(t *testing.T) {
	type args struct {
		grid      []string
		catJump   int
		mouseJump int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]string{"####F", "#C...", "M...."}, 1, 2}, true},
		{"2", args{[]string{"M.C...F"}, 1, 4}, true},
		{"3", args{[]string{"M.C...F"}, 1, 3}, false},
		{"4", args{[]string{"C...#", "...#F", "....#", "M...."}, 2, 5}, false},
		{"5", args{[]string{".M...", "..#..", "#..#.", "C#.#.", "...#F"}, 3, 1}, true},
		{"6", args{[]string{"###.#", "##.F#", "C###.", ".#M.#", "####."}, 2, 2}, false},
		{"7", args{[]string{"........", "F...#C.M", "........"}, 1, 2}, true},
		{"8", args{[]string{"####.##", ".#C#F#.", "######.", "##M.###"}, 3, 6}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMouseWin(tt.args.grid, tt.args.catJump, tt.args.mouseJump); got != tt.want {
				t.Errorf("canMouseWin() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ####.##
// .#C#F#.
// ######.
// ##M.###

package main

import "testing"

func Test_findMinStep(t *testing.T) {
	type args struct {
		board string
		hand  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"WRRBBW", "RB"}, -1},
		{"2", args{"WWRRBBWW", "WRBRW"}, 2},
		{"3", args{"G", "GGGGG"}, 2},
		{"4", args{"RBYYBBRRB", "YRBGB"}, 3},
		{"5", args{"WWBBWBBWW", "BB"}, -1},
		{"6", args{"RRWWRRBBRR", "WB"}, 2},
		{"7", args{"RRYRRYYRYYRRYYRR", "YYRYY"}, 2},
		{"8", args{"RRWWRRW", "WR"}, -1},
		{"9", args{"RRGGBBYYWWRRGGBB", "RGBYW"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinStep(tt.args.board, tt.args.hand); got != tt.want {
				t.Errorf("findMinStep() = %v, want %v", got, tt.want)
			}
		})
	}
}

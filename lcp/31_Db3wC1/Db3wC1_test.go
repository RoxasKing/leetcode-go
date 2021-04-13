package main

import "testing"

func Test_escapeMaze(t *testing.T) {
	type args struct {
		maze [][]string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[][]string{{".#.", "#.."}, {"...", ".#."}, {".##", ".#."}, {"..#", ".#."}}}, true},
		{"2", args{[][]string{{".#.", "..."}, {"...", "..."}}}, false},
		{"3", args{[][]string{{"...", "...", "..."}, {".##", "###", "##."}, {".##", "###", "##."}, {".##", "###", "##."}, {".##", "###", "##."}, {".##", "###", "##."}, {".##", "###", "##."}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := escapeMaze(tt.args.maze); got != tt.want {
				t.Errorf("escapeMaze() = %v, want %v", got, tt.want)
			}
		})
	}
}

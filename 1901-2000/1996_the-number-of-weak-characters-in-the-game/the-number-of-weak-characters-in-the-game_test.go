package main

import "testing"

func Test_numberOfWeakCharacters(t *testing.T) {
	type args struct {
		properties [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{5, 5}, {6, 3}, {3, 6}}}, 0},
		{"2", args{[][]int{{2, 2}, {3, 3}}}, 1},
		{"3", args{[][]int{{1, 5}, {10, 4}, {4, 3}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfWeakCharacters(tt.args.properties); got != tt.want {
				t.Errorf("numberOfWeakCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}

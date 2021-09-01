package main

import "testing"

func Test_minRecSize(t *testing.T) {
	type args struct {
		lines [][]int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1", args{[][]int{{2, 3}, {3, 0}, {4, 1}}}, 48.00000},
		{"2", args{[][]int{{1, 1}, {2, 3}}}, 0.00000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minRecSize(tt.args.lines); got != tt.want {
				t.Errorf("minRecSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

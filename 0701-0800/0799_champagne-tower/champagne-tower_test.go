package main

import "testing"

func Test_champagneTower(t *testing.T) {
	type args struct {
		poured      int
		query_row   int
		query_glass int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1", args{1, 1, 1}, 0.0},
		{"2", args{2, 1, 1}, 0.5},
		{"3", args{100000009, 33, 17}, 1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := champagneTower(tt.args.poured, tt.args.query_row, tt.args.query_glass); got != tt.want {
				t.Errorf("champagneTower() = %v, want %v", got, tt.want)
			}
		})
	}
}

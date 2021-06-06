package main

import "testing"

func Test_computeArea(t *testing.T) {
	type args struct {
		ax1 int
		ay1 int
		ax2 int
		ay2 int
		bx1 int
		by1 int
		bx2 int
		by2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{-3, 0, 3, 4, 0, -1, 9, 2}, 45},
		{"2", args{-2, -2, 2, 2, -2, -2, 2, 2}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeArea(tt.args.ax1, tt.args.ay1, tt.args.ax2, tt.args.ay2, tt.args.bx1, tt.args.by1, tt.args.bx2, tt.args.by2); got != tt.want {
				t.Errorf("computeArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_countBalls(t *testing.T) {
	type args struct {
		lowLimit  int
		highLimit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1, 10}, 2},
		{"2", args{5, 15}, 2},
		{"3", args{19, 28}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBalls(tt.args.lowLimit, tt.args.highLimit); got != tt.want {
				t.Errorf("countBalls() = %v, want %v", got, tt.want)
			}
		})
	}
}

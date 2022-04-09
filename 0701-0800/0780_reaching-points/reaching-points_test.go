package main

import "testing"

func Test_reachingPoints(t *testing.T) {
	type args struct {
		sx int
		sy int
		tx int
		ty int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{1, 1, 3, 5}, true},
		{"2", args{1, 1, 2, 2}, false},
		{"3", args{1, 1, 1, 1}, true},
		{"4", args{9, 10, 9, 19}, true},
		{"5", args{1, 5, 999999996, 5}, true},
		{"6", args{1, 5, 9, 15}, false},
		{"7", args{1, 4, 999999998, 4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reachingPoints(tt.args.sx, tt.args.sy, tt.args.tx, tt.args.ty); got != tt.want {
				t.Errorf("reachingPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

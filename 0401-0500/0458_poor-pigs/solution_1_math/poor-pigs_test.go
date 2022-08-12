package main

import "testing"

func Test_poorPigs(t *testing.T) {
	type args struct {
		buckets       int
		minutesToDie  int
		minutesToTest int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1000, 15, 60}, 5},
		{"2", args{4, 15, 15}, 2},
		{"3", args{4, 15, 30}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := poorPigs(tt.args.buckets, tt.args.minutesToDie, tt.args.minutesToTest); got != tt.want {
				t.Errorf("poorPigs() = %v, want %v", got, tt.want)
			}
		})
	}
}

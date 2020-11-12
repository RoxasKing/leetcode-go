package main

import "testing"

func Test_maxEnvelopes(t *testing.T) {
	type args struct {
		envelopes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{[][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}},
			3,
		},
		{
			"2",
			args{[][]int{{5, 4}, {6, 5}, {6, 7}, {2, 3}}},
			3,
		},
		{
			"3",
			args{[][]int{{5, 4}, {6, 5}, {6, 7}, {2, 3}, {7, 7}}},
			4,
		},
		{
			"4",
			args{[][]int{{1, 3}, {3, 5}, {6, 7}, {6, 8}, {8, 4}, {9, 5}}},
			3,
		},
		{
			"5",
			args{[][]int{{2, 100}, {3, 200}, {4, 300}, {5, 500}, {5, 400}, {5, 250}, {6, 370}, {6, 360}, {7, 380}}},
			5,
		},
		{
			"6",
			args{[][]int{{46, 89}, {50, 53}, {52, 68}, {72, 45}, {77, 81}}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnvelopes(tt.args.envelopes); got != tt.want {
				t.Errorf("maxEnvelopes() = %v, want %v", got, tt.want)
			}
		})
	}
}

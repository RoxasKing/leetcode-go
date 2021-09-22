package main

import "testing"

func Test_maxmiumScore(t *testing.T) {
	type args struct {
		cards []int
		cnt   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 8, 9}, 3}, 18},
		{"2", args{[]int{3, 3, 1}, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxmiumScore(tt.args.cards, tt.args.cnt); got != tt.want {
				t.Errorf("maxmiumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_stoneGameIX(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{2, 1}}, true},
		{"2", args{[]int{2}}, false},
		{"3", args{[]int{5, 1, 2, 4, 3}}, false},
		{"4", args{[]int{2, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGameIX(tt.args.stones); got != tt.want {
				t.Errorf("stoneGameIX() = %v, want %v", got, tt.want)
			}
		})
	}
}

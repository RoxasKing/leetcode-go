package main

import "testing"

func Test_maxCompatibilitySum(t *testing.T) {
	type args struct {
		students [][]int
		mentors  [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 1}}, [][]int{{1, 0, 0}, {0, 0, 1}, {1, 1, 0}}}, 8},
		{"2", args{[][]int{{0, 0}, {0, 0}, {0, 0}}, [][]int{{1, 1}, {1, 1}, {1, 1}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCompatibilitySum(tt.args.students, tt.args.mentors); got != tt.want {
				t.Errorf("maxCompatibilitySum() = %v, want %v", got, tt.want)
			}
		})
	}
}

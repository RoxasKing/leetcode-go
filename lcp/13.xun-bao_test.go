package leetcode

import "testing"

func Test_minimalSteps(t *testing.T) {
	type args struct {
		maze []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]string{"S#O", "M..", "M.T"}}, 16},
		{"2", args{[]string{"S#O", "M.#", "M.T"}}, -1},
		{"3", args{[]string{"S#O", "M.T", "M.T"}}, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimalSteps(tt.args.maze); got != tt.want {
				t.Errorf("minimalSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_halfQuestions(t *testing.T) {
	type args struct {
		questions []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 1, 6, 2}}, 1},
		{"2", args{[]int{1, 5, 1, 3, 4, 5, 2, 5, 3, 3, 8, 6}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := halfQuestions(tt.args.questions); got != tt.want {
				t.Errorf("halfQuestions() = %v, want %v", got, tt.want)
			}
		})
	}
}

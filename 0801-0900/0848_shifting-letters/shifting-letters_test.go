package main

import "testing"

func Test_shiftingLetters(t *testing.T) {
	type args struct {
		s      string
		shifts []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"abc", []int{3, 5, 9}}, "rpl"},
		{"2", args{"aaa", []int{1, 2, 3}}, "gfd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiftingLetters(tt.args.s, tt.args.shifts); got != tt.want {
				t.Errorf("shiftingLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

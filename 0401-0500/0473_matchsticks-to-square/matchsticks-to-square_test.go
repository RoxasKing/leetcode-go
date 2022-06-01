package main

import "testing"

func Test_makesquare(t *testing.T) {
	type args struct {
		matchsticks []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{1, 1, 2, 2, 2}}, true},
		{"2", args{[]int{3, 3, 3, 3, 4}}, false},
		{"3", args{[]int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}}, true},
		{"4", args{[]int{5, 5, 5, 5, 16, 4, 4, 4, 4, 4, 3, 3, 3, 3, 4}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makesquare(tt.args.matchsticks); got != tt.want {
				t.Errorf("makesquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

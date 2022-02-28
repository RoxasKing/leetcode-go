package main

import "testing"

func Test_optimalDivision(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{[]int{1000, 100, 10, 2}}, "1000/(100/10/2)"},
		{"2", args{[]int{2, 3, 4}}, "2/(3/4)"},
		{"3", args{[]int{2}}, "2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := optimalDivision(tt.args.nums); got != tt.want {
				t.Errorf("optimalDivision() = %v, want %v", got, tt.want)
			}
		})
	}
}

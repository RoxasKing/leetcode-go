package main

import "testing"

func Test_validMountainArray(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{2, 1}}, false},
		{"2", args{[]int{3, 5, 5}}, false},
		{"3", args{[]int{0, 3, 2, 1}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArray(tt.args.A); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

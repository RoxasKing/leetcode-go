package main

import "testing"

func Test_maximumNumber(t *testing.T) {
	type args struct {
		num    string
		change []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"132", []int{9, 8, 5, 0, 3, 6, 4, 2, 6, 8}}, "832"},
		{"2", args{"021", []int{9, 4, 3, 5, 7, 2, 1, 9, 0, 6}}, "934"},
		{"3", args{"5", []int{1, 4, 7, 5, 3, 2, 5, 6, 9, 4}}, "5"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumNumber(tt.args.num, tt.args.change); got != tt.want {
				t.Errorf("maximumNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

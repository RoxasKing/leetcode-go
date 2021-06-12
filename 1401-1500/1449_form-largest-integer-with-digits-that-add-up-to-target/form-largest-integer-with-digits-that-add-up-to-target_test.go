package main

import "testing"

func Test_largestNumber(t *testing.T) {
	type args struct {
		cost   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{[]int{4, 3, 2, 5, 6, 7, 2, 5, 5}, 9}, "7772"},
		{"2", args{[]int{7, 6, 5, 5, 5, 6, 8, 7, 8}, 12}, "85"},
		{"3", args{[]int{2, 4, 6, 2, 4, 6, 4, 4, 4}, 5}, "0"},
		{"4", args{[]int{6, 10, 15, 40, 40, 40, 40, 40, 40}, 47}, "32211"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.args.cost, tt.args.target); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

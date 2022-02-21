package main

import "testing"

func Test_pancakeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{[]int{3, 2, 4, 1}}},
		{"2", args{[]int{1, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pancakeSort(tt.args.arr); len(got) > 10*len(tt.args.arr) {
				t.Errorf("len(pancakeSort(arr)) > 10 * len(arr)")
			}
		})
	}
}

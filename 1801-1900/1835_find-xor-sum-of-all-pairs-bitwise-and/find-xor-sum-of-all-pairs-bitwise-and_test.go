package main

import "testing"

func Test_getXORSum(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3}, []int{6, 5}}, 0},
		{"2", args{[]int{12}, []int{4}}, 4},
		{"3", args{[]int{1, 2, 2}, []int{6, 5}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getXORSum(tt.args.arr1, tt.args.arr2); got != tt.want {
				t.Errorf("getXORSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

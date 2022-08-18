package main

import "testing"

func Test_minSetSize(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}}, 2},
		{"2", args{[]int{7, 7, 7, 7, 7, 7}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSetSize(tt.args.arr); got != tt.want {
				t.Errorf("minSetSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

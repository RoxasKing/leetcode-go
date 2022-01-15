package main

import "testing"

func Test_minJumps(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}}, 3},
		{"2", args{[]int{7}}, 0},
		{"3", args{[]int{7, 6, 9, 6, 9, 6, 9, 7}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minJumps(tt.args.arr); got != tt.want {
				t.Errorf("minJumps() = %v, want %v", got, tt.want)
			}
		})
	}
}

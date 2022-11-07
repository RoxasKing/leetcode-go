package main

import "testing"

func Test_earliestFullBloom(t *testing.T) {
	type args struct {
		plantTime []int
		growTime  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 4, 3}, []int{2, 3, 1}}, 9},
		{"2", args{[]int{1, 2, 3, 2}, []int{2, 1, 2, 1}}, 9},
		{"3", args{[]int{1}, []int{1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := earliestFullBloom(tt.args.plantTime, tt.args.growTime); got != tt.want {
				t.Errorf("earliestFullBloom() = %v, want %v", got, tt.want)
			}
		})
	}
}

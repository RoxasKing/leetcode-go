package main

import "testing"

func Test_storeWater(t *testing.T) {
	b := []int{}
	v := []int{}
	for i := 0; i < 100; i++ {
		b = append(b, 0)
		v = append(v, 1e4)
	}
	type args struct {
		bucket []int
		vat    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 3}, []int{6, 8}}, 4},
		{"2", args{[]int{9, 0, 1}, []int{0, 2, 2}}, 3},
		{"3", args{[]int{1, 2}, []int{100, 50}}, 22},
		{"4", args{[]int{1}, []int{1e4}}, 199},
		{"5", args{b, v}, 2000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := storeWater(tt.args.bucket, tt.args.vat); got != tt.want {
				t.Errorf("storeWater() = %v, want %v", got, tt.want)
			}
		})
	}
}

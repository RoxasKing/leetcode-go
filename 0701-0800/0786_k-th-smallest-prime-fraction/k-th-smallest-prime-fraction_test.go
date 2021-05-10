package main

import (
	"reflect"
	"testing"
)

func Test_kthSmallestPrimeFraction(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 2, 3, 5}, 3}, []int{2, 5}},
		{"2", args{[]int{1, 7}, 1}, []int{1, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallestPrimeFraction(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kthSmallestPrimeFraction() = %v, want %v", got, tt.want)
			}
		})
	}
}

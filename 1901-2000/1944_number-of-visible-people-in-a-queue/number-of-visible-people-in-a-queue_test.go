package main

import (
	"reflect"
	"testing"
)

func Test_canSeePersonsCount(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{10, 6, 8, 5, 11, 9}}, []int{3, 1, 2, 1, 1, 0}},
		{"2", args{[]int{5, 1, 2, 3, 10}}, []int{4, 1, 1, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canSeePersonsCount(tt.args.heights); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("canSeePersonsCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

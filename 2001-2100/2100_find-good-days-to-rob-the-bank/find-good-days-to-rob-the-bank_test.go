package main

import (
	"reflect"
	"testing"
)

func Test_goodDaysToRobBank(t *testing.T) {
	type args struct {
		security []int
		time     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{5, 3, 3, 3, 5, 6, 2}, 2}, []int{2, 3}},
		{"2", args{[]int{1, 1, 1, 1, 1}, 0}, []int{0, 1, 2, 3, 4}},
		{"3", args{[]int{1, 2, 3, 4, 5, 6}, 2}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goodDaysToRobBank(tt.args.security, tt.args.time); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("goodDaysToRobBank() = %v, want %v", got, tt.want)
			}
		})
	}
}

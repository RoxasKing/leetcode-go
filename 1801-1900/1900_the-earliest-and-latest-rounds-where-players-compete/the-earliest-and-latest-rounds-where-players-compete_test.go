package main

import (
	"reflect"
	"testing"
)

func Test_earliestAndLatest(t *testing.T) {
	type args struct {
		n            int
		firstPlayer  int
		secondPlayer int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{11, 2, 4}, []int{3, 4}},
		{"2", args{5, 1, 5}, []int{1, 1}},
		{"3", args{28, 13, 17}, []int{3, 5}},
		{"4", args{27, 12, 15}, []int{3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := earliestAndLatest(tt.args.n, tt.args.firstPlayer, tt.args.secondPlayer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("earliestAndLatest() = %v, want %v", got, tt.want)
			}
		})
	}
}

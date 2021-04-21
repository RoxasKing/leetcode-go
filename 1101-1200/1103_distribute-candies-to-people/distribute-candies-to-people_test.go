package main

import (
	"reflect"
	"testing"
)

func Test_distributeCandies(t *testing.T) {
	type args struct {
		candies   int
		numPeople int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{0, 5}, []int{0, 0, 0, 0, 0}},
		{"2", args{2, 5}, []int{1, 1, 0, 0, 0}},
		{"3", args{7, 4}, []int{1, 2, 3, 1}},
		{"5", args{10, 3}, []int{5, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCandies(tt.args.candies, tt.args.numPeople); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}

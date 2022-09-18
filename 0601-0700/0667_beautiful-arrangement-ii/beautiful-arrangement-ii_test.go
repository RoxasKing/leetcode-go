package main

import (
	"reflect"
	"testing"
)

func Test_constructArray(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{3, 1}, []int{1, 2, 3}},
		{"2", args{3, 2}, []int{1, 3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructArray(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

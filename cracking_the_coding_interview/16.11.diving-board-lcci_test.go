package main

import (
	"reflect"
	"testing"
)

func Test_divingBoard(t *testing.T) {
	type args struct {
		shorter int
		longer  int
		k       int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{1, 2, 3}, []int{3, 4, 5, 6}},
		{"2", args{1, 1, 3}, []int{3}},
		{"3", args{1, 1, 0}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divingBoard(tt.args.shorter, tt.args.longer, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divingBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

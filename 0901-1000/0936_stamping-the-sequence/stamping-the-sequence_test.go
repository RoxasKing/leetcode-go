package main

import (
	"reflect"
	"testing"
)

func Test_movesToStamp(t *testing.T) {
	type args struct {
		stamp  string
		target string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{"abc", "ababc"}, []int{0, 2}},
		{"2", args{"abca", "aabcaca"}, []int{0, 3, 1}},
		{"3", args{"oz", "ooozz"}, []int{3, 0, 1, 2}},
		{"4", args{"cab", "cabbb"}, []int{2, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movesToStamp(tt.args.stamp, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("movesToStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

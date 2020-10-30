package main

import (
	"testing"
)

func Test_uniqueOccurrences(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{1, 2, 2, 1, 1, 3}}, true},
		{"2", args{[]int{1, 2}}, false},
		{"3", args{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueOccurrences(tt.args.arr); got != tt.want {
				t.Errorf("uniqueOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uniqueOccurrences2(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{1, 2, 2, 1, 1, 3}}, true},
		{"2", args{[]int{1, 2}}, false},
		{"3", args{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueOccurrences2(tt.args.arr); got != tt.want {
				t.Errorf("uniqueOccurrences2() = %v, want %v", got, tt.want)
			}
		})
	}
}

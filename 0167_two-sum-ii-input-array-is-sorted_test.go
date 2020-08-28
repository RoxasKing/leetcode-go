package main

import (
	"reflect"
	"testing"
)

func Test_twoSumII(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{2, 7, 11, 15}, 9}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumII(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSumII2(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{2, 7, 11, 15}, 9}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumII2(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumII2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSumII3(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{2, 7, 11, 15}, 9}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumII3(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumII3() = %v, want %v", got, tt.want)
			}
		})
	}
}

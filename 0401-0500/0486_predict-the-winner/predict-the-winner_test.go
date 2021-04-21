package main

import (
	"testing"
)

func TestPredictTheWinner(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{1, 5, 2}}, false},
		{"2", args{[]int{1, 5, 233, 7}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PredictTheWinner(tt.args.nums); got != tt.want {
				t.Errorf("PredictTheWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPredictTheWinner2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{1, 5, 2}}, false},
		{"2", args{[]int{1, 5, 233, 7}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PredictTheWinner2(tt.args.nums); got != tt.want {
				t.Errorf("PredictTheWinner2() = %v, want %v", got, tt.want)
			}
		})
	}
}

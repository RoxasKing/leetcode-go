package main

import (
	"testing"
)

func Test_translateNum(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{12258}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateNum(tt.args.num); got != tt.want {
				t.Errorf("translateNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_translateNum2(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{12258}, 5},
		{"2", args{18822}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateNum2(tt.args.num); got != tt.want {
				t.Errorf("translateNum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

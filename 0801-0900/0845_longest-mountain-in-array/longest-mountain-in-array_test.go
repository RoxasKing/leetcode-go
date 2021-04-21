package main

import (
	"testing"
)

func Test_longestMountain(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{"1", args{[]int{2, 1, 4, 7, 3, 2, 5}}, 5},
		{"2", args{[]int{2, 2, 2}}, 0},
		{"3", args{[]int{2, 1, 4, 7, 4, 3, 2, 5}}, 6},
		{"4", args{[]int{2, 1, 4, 7, 4, 3, 2, 5, 6, 9, 8, 3, 1}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestMountain(tt.args.a); gotAns != tt.wantAns {
				t.Errorf("longestMountain() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

func Test_longestMountain2(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{"1", args{[]int{2, 1, 4, 7, 3, 2, 5}}, 5},
		{"2", args{[]int{2, 2, 2}}, 0},
		{"3", args{[]int{2, 1, 4, 7, 4, 3, 2, 5}}, 6},
		{"4", args{[]int{2, 1, 4, 7, 4, 3, 2, 5, 6, 9, 8, 3, 1}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestMountain2(tt.args.a); gotAns != tt.wantAns {
				t.Errorf("longestMountain2() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

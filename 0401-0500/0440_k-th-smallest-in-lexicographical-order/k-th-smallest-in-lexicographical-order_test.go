package main

import "testing"

func Test_findKthNumber(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{13, 2}, 10},
		{"2", args{1, 1}, 1},
		{"3", args{1356, 36}, 103},
		{"4", args{10000, 10000}, 9999},
		{"5", args{200, 111}, 199},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthNumber(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("findKthNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

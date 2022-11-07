package main

import "testing"

func Test_maximum69Number(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{9669}, 9969},
		{"2", args{9996}, 9999},
		{"3", args{9999}, 9999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum69Number(tt.args.num); got != tt.want {
				t.Errorf("maximum69Number() = %v, want %v", got, tt.want)
			}
		})
	}
}

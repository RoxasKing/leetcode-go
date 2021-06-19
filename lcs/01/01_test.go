package main

import "testing"

func Test_leastMinutes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{2}, 2},
		{"2", args{4}, 3},
		{"3", args{3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastMinutes(tt.args.n); got != tt.want {
				t.Errorf("leastMinutes() = %v, want %v", got, tt.want)
			}
		})
	}
}

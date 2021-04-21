package main

import "testing"

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{0b00000000000000000000000000001011}, 3},
		{"2", args{0b00000000000000000000000010000000}, 1},
		{"3", args{0b11111111111111111111111111111101}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

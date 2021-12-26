package main

import "testing"

func Test_atMostNGivenDigitSet(t *testing.T) {
	type args struct {
		digits []string
		n      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]string{"1", "3", "5", "7"}, 100}, 20},
		{"2", args{[]string{"1", "4", "9"}, 1000000000}, 29523},
		{"3", args{[]string{"7"}, 8}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := atMostNGivenDigitSet(tt.args.digits, tt.args.n); got != tt.want {
				t.Errorf("atMostNGivenDigitSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

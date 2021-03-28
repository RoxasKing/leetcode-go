package main

import "testing"

func Test_maxNiceDivisors(t *testing.T) {
	type args struct {
		primeFactors int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5}, 6},
		{"2", args{8}, 18},
		{"3", args{1000}, 620946522},
		{"4", args{10000}, 515481589},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNiceDivisors(tt.args.primeFactors); got != tt.want {
				t.Errorf("maxNiceDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}

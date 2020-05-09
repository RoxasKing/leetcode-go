package codinginterviews

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1", args{2.0, 10}, 1024.0},
		{"2", args{2.1, 3}, 9.261},
		{"3", args{2.0, -2}, 0.25},
		{"4", args{0.00001, 2147483647}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}

package codinginterviews

import (
	"math"
	"testing"
)

func Test_twoSumII(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{"1", args{1}, []float64{0.16667, 0.16667, 0.16667, 0.16667, 0.16667, 0.16667}},
		{"2", args{2}, []float64{0.02778, 0.05556, 0.08333, 0.11111, 0.13889, 0.16667, 0.13889, 0.11111, 0.08333, 0.05556, 0.02778}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSumII(tt.args.n)
			for i := range got {
				if math.Abs(got[i]-tt.want[i]) > 1e-5 {
					t.Errorf("twoSumII() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}

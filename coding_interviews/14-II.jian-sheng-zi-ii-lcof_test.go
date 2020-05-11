package codinginterviews

import "testing"

func Test_cuttingRopeII(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{2}, 1},
		{"2", args{10}, 36},
		{"3", args{120}, 953271190},
		{"4", args{127}, 439254203},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cuttingRopeII(tt.args.n); got != tt.want {
				t.Errorf("cuttingRopeII() = %v, want %v", got, tt.want)
			}
		})
	}
}

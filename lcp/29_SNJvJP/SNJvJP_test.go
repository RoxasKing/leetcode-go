package main

import "testing"

func Test_orchestraLayout(t *testing.T) {
	type args struct {
		num  int
		xPos int
		yPos int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{3, 0, 2}, 3},
		{"2", args{4, 1, 2}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orchestraLayout(tt.args.num, tt.args.xPos, tt.args.yPos); got != tt.want {
				t.Errorf("orchestraLayout() = %v, want %v", got, tt.want)
			}
		})
	}
}

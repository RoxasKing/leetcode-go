package main

import "testing"

func Test_minDeletionSize(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]string{"cba", "daf", "ghi"}}, 1},
		{"2", args{[]string{"a", "b"}}, 0},
		{"3", args{[]string{"zyx", "wvu", "tsr"}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDeletionSize(tt.args.strs); got != tt.want {
				t.Errorf("minDeletionSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

package codinginterviews

import (
	"testing"
)

func Test_verifyPostorder(t *testing.T) {
	type args struct {
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{1, 6, 3, 2, 5}}, false},
		{"2", args{[]int{1, 3, 2, 6, 5}}, true},
		{"3", args{[]int{1, 2, 5, 10, 6, 9, 4, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verifyPostorder(tt.args.postorder); got != tt.want {
				t.Errorf("verifyPostorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

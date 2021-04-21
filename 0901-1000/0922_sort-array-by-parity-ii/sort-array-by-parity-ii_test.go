package main

import (
	"testing"
)

func Test_sortArrayByParityII(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{[]int{4, 2, 5, 7}}},
		{"2", args{[]int{4, 2, 5, 7, 3, 9, 6, 8}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParityII(tt.args.A); !isValid(got) {
				t.Errorf("sortArrayByParityII() test failed!")
			}
		})
	}
}

func isValid(arr []int) bool {
	for i := range arr {
		if i&1 == 1 && arr[i]&1 == 0 || i&1 == 0 && arr[i]&1 == 1 {
			return false
		}
	}
	return true
}

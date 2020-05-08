package codinginterviews

import (
	"reflect"
	"testing"
)

func Test_singleNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{4, 1, 4, 6}}, []int{6, 1}},
		{"2", args{[]int{1, 2, 10, 4, 1, 4, 3, 3}}, []int{2, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{1}}, []int{1}},
		{"", args{[]int{2, 1}}, []int{1, 2}},
		{"", args{[]int{7, 10, 8, 9, 3, 2, 1, 11, 6, 5, 4, 12, 14, 13, 0}},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.array)
			if !reflect.DeepEqual(tt.args.array, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", tt.args.array, tt.want)
			}
		})
	}
}

package coding_interview

import (
	"reflect"
	"testing"
)

func Test_getLeastNumbers(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{3, 2, 1}, 2}, []int{2, 1}},
		{"", args{[]int{0, 1, 2, 1}, 1}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLeastNumbers(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLeastNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLeastNumbers2(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{3, 2, 1}, 2}, []int{1, 2}},
		{"", args{[]int{0, 1, 2, 1}, 1}, []int{0}},
		{"", args{[]int{0, 0, 2, 3, 2, 1, 1, 2, 0, 4}, 10}, []int{0, 0, 0, 1, 1, 2, 2, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLeastNumbers2(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLeastNumbers2() = %v, want %v", got, tt.want)
			}
		})
	}
}

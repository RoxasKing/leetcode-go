package codinginterviews

import (
	"reflect"
	"testing"
)

func Test_constructArr(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 2, 3, 4, 5}}, []int{120, 60, 40, 30, 24}},
		{"2", args{[]int{0, 2, 3, 4, 5}}, []int{120, 0, 0, 0, 0}},
		{"3", args{[]int{0, 0, 3, 4, 5}}, []int{0, 0, 0, 0, 0}},
		{"4", args{nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructArr(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructArr() = %v, want %v", got, tt.want)
			}
		})
	}
}

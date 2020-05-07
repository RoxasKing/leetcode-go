package crackingthecodingintervew

import (
	"reflect"
	"testing"
)

func Test_intersection(t *testing.T) {
	type args struct {
		start1 []int
		end1   []int
		start2 []int
		end2   []int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{"1", args{[]int{0, 0}, []int{1, 0}, []int{1, 1}, []int{0, -1}}, []float64{0.5, 0}},
		{"2", args{[]int{0, 0}, []int{3, 3}, []int{1, 1}, []int{2, 2}}, []float64{1, 1}},
		{"3", args{[]int{0, 0}, []int{1, 1}, []int{1, 0}, []int{2, 1}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.args.start1, tt.args.end1, tt.args.start2, tt.args.end2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

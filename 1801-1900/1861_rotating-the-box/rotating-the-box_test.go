package main

import (
	"reflect"
	"testing"
)

func Test_rotateTheBox(t *testing.T) {
	type args struct {
		box [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{"1", args{[][]byte{{'#', '#', '*', '.', '*', '.'}, {'#', '#', '#', '*', '.', '.'}, {'#', '#', '#', '.', '#', '.'}}}, [][]byte{{'.', '#', '#'}, {'.', '#', '#'}, {'#', '#', '*'}, {'#', '*', '.'}, {'#', '.', '*'}, {'#', '.', '.'}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateTheBox(tt.args.box); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateTheBox() = %v, want %v", got, tt.want)
			}
		})
	}
}

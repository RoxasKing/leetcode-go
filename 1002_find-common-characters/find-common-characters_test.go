package main

import (
	"reflect"
	"testing"
)

func Test_commonChars(t *testing.T) {
	type args struct {
		A []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{[]string{"bella", "label", "roller"}}, []string{"e", "l", "l"}},
		{"2", args{[]string{"cool", "lock", "cook"}}, []string{"c", "o"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commonChars(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commonChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

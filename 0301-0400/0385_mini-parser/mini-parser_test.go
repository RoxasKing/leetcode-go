package main

import (
	"reflect"
	"testing"
)

func Test_deserialize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *NestedInteger
	}{
		{
			"1",
			args{"324"},
			&NestedInteger{
				isInt: true,
				value:     324,
			},
		},
		{
			"2",
			args{"[123,[456,[789]]]"},
			&NestedInteger{
				value: -1 << 31,
				nlist: []*NestedInteger{
					{
						isInt: true,
						value:     123,
					},
					{
						value: -1 << 31,
						nlist: []*NestedInteger{
							{
								isInt: true,
								value:     456,
							},
							{
								value: -1 << 31,
								nlist: []*NestedInteger{
									{
										isInt: true,
										value:     789,
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deserialize(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_findItinerary(t *testing.T) {
	type args struct {
		tickets [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"1",
			args{
				[][]string{
					{"MUC", "LHR"},
					{"JFK", "MUC"},
					{"SFO", "SJC"},
					{"LHR", "SFO"},
				},
			},
			[]string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		},
		{
			"2",
			args{
				[][]string{
					{"JFK", "SFO"},
					{"JFK", "ATL"},
					{"SFO", "ATL"},
					{"ATL", "JFK"},
					{"ATL", "SFO"},
				},
			},
			[]string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
		},
		{
			"3",
			args{
				[][]string{
					{"JFK", "KUL"},
					{"JFK", "NRT"},
					{"NRT", "JFK"},
				},
			},
			[]string{"JFK", "NRT", "JFK", "KUL"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findItinerary(tt.args.tickets); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findItinerary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findItinerary2(t *testing.T) {
	type args struct {
		tickets [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"1",
			args{
				[][]string{
					{"MUC", "LHR"},
					{"JFK", "MUC"},
					{"SFO", "SJC"},
					{"LHR", "SFO"},
				},
			},
			[]string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		},
		{
			"2",
			args{
				[][]string{
					{"JFK", "SFO"},
					{"JFK", "ATL"},
					{"SFO", "ATL"},
					{"ATL", "JFK"},
					{"ATL", "SFO"},
				},
			},
			[]string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
		},
		{
			"3",
			args{
				[][]string{
					{"JFK", "KUL"},
					{"JFK", "NRT"},
					{"NRT", "JFK"},
				},
			},
			[]string{"JFK", "NRT", "JFK", "KUL"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findItinerary2(tt.args.tickets); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findItinerary2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package day05

import (
	"reflect"
	"testing"
)

func TestAlmanac_AddSeeds(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []int64
	}{
		{
			name: "success",
			s:    "seeds: 79 14 55 13",
			want: []int64{79, 14, 55, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Almanac{}
			if got := a.AddSeeds(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddSeeds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve(t *testing.T) {
	tests := []struct {
		name string
		f    string
		want int64
	}{
		{
			name: "success",
			f:    "input.txt",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.f); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

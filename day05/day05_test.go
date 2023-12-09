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

package day04

import (
	"reflect"
	"testing"
)

func Test_parseNumbers(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{n: "96 13 61 89 70 80 93 57  9 28"},
			want: []int64{96, 13, 61, 89, 70, 80, 93, 57, 9, 28},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseNumbers(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
		{name: "success", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

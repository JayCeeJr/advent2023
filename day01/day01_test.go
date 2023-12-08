package day01

import "testing"

func Test_solve(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{
			name: "success",
			want: 55971,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(); got != tt.want {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_solve2(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{
			name: "success",
			want: 55200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve2(); got != tt.want {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

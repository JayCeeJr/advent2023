package day01

import "testing"

func Test_readFile(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "success",
			args: args{name: "input.txt"},
			want: 55971,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readFile(tt.args.name); got != tt.want {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

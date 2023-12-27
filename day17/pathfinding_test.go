package main

import "testing"

func Test_optimalPath(t *testing.T) {
	type args struct {
		input [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "simple case - no movement restrictions",
			args: args{},
			want: 78,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			part1("test.txt")
			if got := optimalPath(); got != tt.want {
				t.Errorf("optimalPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

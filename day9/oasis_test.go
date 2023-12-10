package main

import "testing"

func Test_predictNextValue(t *testing.T) {
	type args struct {
		valueHistory []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple case - 2 sequences",
			args: args{valueHistory: []int{0, 3, 6, 9, 12, 15}},
			want: 18,
		},
		{
			name: "3 sequences",
			args: args{valueHistory: []int{1, 3, 6, 10, 15, 21}},
			want: 28,
		},
		{
			name: "4 sequences",
			args: args{valueHistory: []int{10, 13, 16, 21, 30, 45}},
			want: 68,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := predictNextValue(tt.args.valueHistory); got != tt.want {
				t.Errorf("predictNextValue() = %v, want %v", got, tt.want)
			}
			sequencesForHistory = [][]int{}
		})
	}
}

func Test_predictPreviousValue(t *testing.T) {
	type args struct {
		valueHistory []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	name: "simple case - 2 sequences",
		//	args: args{valueHistory: []int{0, 3, 6, 9, 12, 15}},
		//	want: 18,
		//},
		//{
		//	name: "3 sequences",
		//	args: args{valueHistory: []int{1, 3, 6, 10, 15, 21}},
		//	want: 28,
		//},
		{
			name: "4 sequences",
			args: args{valueHistory: []int{10, 13, 16, 21, 30, 45}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := predictPreviousValue(tt.args.valueHistory); got != tt.want {
				t.Errorf("predictPrevious() = %v, want %v", got, tt.want)
			}
			sequencesForHistory = [][]int{}
		})
	}
}

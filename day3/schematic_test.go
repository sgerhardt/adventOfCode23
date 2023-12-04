package main

import (
	"reflect"
	"testing"
)

func Test_findPartNumbers(t *testing.T) {
	type args struct {
		in0 [][]rune
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "no part numbers", args: args{[][]rune{{'.', '1', '4'}}}, want: nil},
		{name: "no part numbers", args: args{[][]rune{{'.', '1', '4', '.', '4'}}}, want: nil},
		{name: "single part horizontal", args: args{[][]rune{{'*', '1', '4'}}}, want: []int{14}},
		{name: "single part horizontal", args: args{[][]rune{{'1', '6', '#'}}}, want: []int{16}},
		{name: "single part horizontal", args: args{[][]rune{{'1', '6', '#', '.', '4'}}}, want: []int{16}},
		{name: "single part horizontal", args: args{[][]rune{{'.', '1', '7', '#'}}}, want: []int{17}},
		{name: "multiple parts horizontal", args: args{[][]rune{{'.', '1', '7', '#', '.', '$', '5'}}}, want: []int{17, 5}},
		{name: "single part vertical", args: args{[][]rune{
			{'.', '1', '8', '.'},
			{'.', '%', '.', '.'},
		}}, want: []int{18}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPartNumbers(tt.args.in0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPartNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

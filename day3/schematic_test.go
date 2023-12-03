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

func Test_findGearRatios(t *testing.T) {
	type args struct {
		s [][]rune
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "no gear ratios", args: args{[][]rune{{'.', '1', '4'}}}, want: nil},
		{name: "simple gear ratio - single row", args: args{[][]rune{{'2', '*', '3'}}}, want: []int{6}},
		{name: "simple gear ratio - multi row", args: args{[][]rune{
			{'2', '*', '.'},
			{'3', '.', '.'}}},
			want: []int{6}},
		{name: "multi digit gear ratio - single row", args: args{[][]rune{{'2', '*', '2', '0'}}}, want: []int{40}},
		{name: "multi digit gear ratio - single row", args: args{[][]rune{{'1', '0', '*', '1', '0'}}}, want: []int{100}},
		{name: "multi digit gear ratio - single row", args: args{[][]rune{{'1', '0', '*', '1', '0'}}}, want: []int{100}},
		{name: "multi digit gear ratio - multi row", args: args{[][]rune{
			{'3', '0', '.', '.', '.'},
			{'.', '.', '*', '1', '0'}}}, want: []int{300}},
		{name: "multi digit gear ratio - example", args: args{[][]rune{
			{'4', '6', '7', '.', '.', '1', '1', '4', '.', '.'},
			{'.', '.', '.', '*', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '3', '5', '.', '.', '6', '3', '3', '.'},
		}}, want: []int{16345}},
		{name: "too many numbers for a gear", args: args{[][]rune{
			{'3', '.', '.'},
			{'4', '*', '1'},
		}}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findGearRatios(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findGearRatios() = %v, want %v", got, tt.want)
			}
		})
	}
}

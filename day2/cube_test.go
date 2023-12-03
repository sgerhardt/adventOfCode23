package main

import "testing"

func Test_isGamePossible(t *testing.T) {
	type args struct {
		bagColors
		gameColors
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "simple", args: args{bagColors{1, 1, 1}, gameColors{1, 1, 1}}, want: true},
		{name: "game not possible", args: args{bagColors{8, 6, 20}, gameColors{12, 14, 13}}, want: false},
		{name: "game not possible with one less", args: args{bagColors{4, 6, 2}, gameColors{12, 14, 13}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isGamePossible(tt.args.bagColors, tt.args.gameColors); got != tt.want {
				t.Errorf("isGamePossible() = %v, want %v", got, tt.want)
			}
		})
	}
}

type setOfBagColors []bagColors

func Test_FindMinSetOfCubesToMakeGamePossible(t *testing.T) {
	type args struct {
		setOfBagColors
		gameColors
		minSetToMakeGamePossible
	}
	tests := []struct {
		name string
		args args
		want bagColors
	}{
		{name: "simple", args: args{[]bagColors{{3, 4, 5}, {1, 4, 5}, {3, 4, 2}}, gameColors{2, 5, 7},
			minSetToMakeGamePossible{
				red:   1,
				blue:  4,
				green: 2,
			}}, want: bagColors{
			red:   3,
			blue:  4,
			green: 5,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinSetToMakePossible(tt.args.setOfBagColors); got != tt.want {
				t.Errorf("isGamePossible() = %v, want %v", got, tt.want)
			}
		})
	}
}

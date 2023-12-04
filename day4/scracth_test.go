package main

import (
	"reflect"
	"testing"
)

func Test_calcWinnings(t *testing.T) {
	type args struct {
		winningNumbers map[string]struct{}
		cardNumbers    map[string]struct{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{name: "no winners", args: args{winningNumbers: map[string]struct{}{}, cardNumbers: map[string]struct{}{"1": struct{}{}}}, want: 0},
		{name: "4 winners", args: args{
			winningNumbers: map[string]struct{}{"41": {}, "48": {}, "83": {}, "86": {}, "17": {}},

			cardNumbers: map[string]struct{}{"48": {}, "83": {}, "86": {}, "17": {}, "9": {}, "1": {}, "2": {}, "3": {}}},
			want: 8},
		{name: "2 winners", args: args{
			winningNumbers: map[string]struct{}{"1": {}, "21": {}, "53": {}, "59": {}, "44": {}},

			cardNumbers: map[string]struct{}{"69": {}, "82": {}, "63": {}, "72": {}, "16": {}, "21": {}, "14": {}, "1": {}}},
			want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcWinnings(tt.args.winningNumbers, tt.args.cardNumbers); got != tt.want {
				t.Errorf("calcWinnings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getScratchcardsForWin(t *testing.T) {
	type args struct {
		card           scratchcard
		winningNumbers map[string]struct{}
		cardNumbers    map[string]struct{}
	}
	tests := []struct {
		name string
		args args
		want []scratchcard
	}{
		{name: "no winners", args: args{card: scratchcard{
			gameNumber:     0,
			winningNumbers: nil,
			cardNumbers:    nil,
		}, winningNumbers: map[string]struct{}{}, cardNumbers: map[string]struct{}{"1": {}}}, want: nil},
		{name: "4 winners", args: args{card: scratchcard{
			gameNumber:     1,
			winningNumbers: map[string]struct{}{"41": {}, "48": {}, "83": {}, "86": {}, "17": {}},
			cardNumbers:    map[string]struct{}{"48": {}, "83": {}, "86": {}, "17": {}, "9": {}, "1": {}, "2": {}, "3": {}}},
		},
			want: []scratchcard{
				{
					gameNumber: 2,
				},
				{
					gameNumber: 3,
				},
				{
					gameNumber: 4,
				},
				{
					gameNumber: 5,
				},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getScratchcardsForWin(tt.args.card); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getScratchcardsForWin() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_countNumWaysToWin(t *testing.T) {
	type args struct {
		raceDuration int
		record       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple race 1",
			args: args{
				raceDuration: 7,
				record:       9,
			},
			want: 4,
		},
		{
			name: "simple race 2",
			args: args{
				raceDuration: 15,
				record:       40,
			},
			want: 8,
		},
		{
			name: "simple race 3",
			args: args{
				raceDuration: 30,
				record:       200,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNumWaysToWin(tt.args.raceDuration, tt.args.record); got != tt.want {
				t.Errorf("countNumWaysToWin() = %v, want %v", got, tt.want)
			}
		})
	}
}

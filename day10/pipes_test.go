package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestPipeDistance(t *testing.T) {

}

func Test_pipeDistance(t *testing.T) {
	type args struct {
		fileName string
		input    func(fileName string) [][]rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple case",
			args: args{
				fileName: "test.txt",
				input: func(fileName string) [][]rune {
					file, err := os.Open(fileName)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
						os.Exit(1)
					}
					defer file.Close()
					scanner := bufio.NewScanner(file)

					var input [][]rune
					idx := 0
					for scanner.Scan() {
						input = append(input, []rune{})
						line := scanner.Text()
						for _, char := range line {
							input[idx] = append(input[idx], char)
						}
						idx++
					}
					return input
				},
			},
			want: 4,
		},
		{
			name: "case with non main loop pipes",
			args: args{
				fileName: "test2.txt",
				input: func(fileName string) [][]rune {
					file, err := os.Open(fileName)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
						os.Exit(1)
					}
					defer file.Close()
					scanner := bufio.NewScanner(file)

					var input [][]rune
					idx := 0
					for scanner.Scan() {
						input = append(input, []rune{})
						line := scanner.Text()
						for _, char := range line {
							input[idx] = append(input[idx], char)
						}
						idx++
					}
					return input
				},
			},
			want: 4,
		},
		{
			name: "more complex case with non main loop pipes",
			args: args{
				fileName: "test3.txt",
				input: func(fileName string) [][]rune {
					file, err := os.Open(fileName)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
						os.Exit(1)
					}
					defer file.Close()
					scanner := bufio.NewScanner(file)

					var input [][]rune
					idx := 0
					for scanner.Scan() {
						input = append(input, []rune{})
						line := scanner.Text()
						for _, char := range line {
							input[idx] = append(input[idx], char)
						}
						idx++
					}
					return input
				},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pipeFarthestPoint(tt.args.input(tt.args.fileName)); got != tt.want {
				t.Errorf("pipeFarthestPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

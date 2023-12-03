package main

import "testing"

func TestCalcSum(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"simple ", args{"1arwwea9"}, 19},
		{"simple single num", args{"zsr1afg"}, 11},
		{"spelt single num", args{"zsroneafg"}, 11},
		{"simple 2", args{"1arw4wea9"}, 19},
		{"simple 3", args{"a1rw4wea9"}, 19},
		{"simple 3", args{"a1rw4we5a9b"}, 19},
		{"spelt", args{"one two three four five six seven eight nine"}, 19},
		{"spelt single num", args{"deronefasf"}, 11},
		{"spelt", args{"oneight"}, 18},
		{"spelt overlap", args{"4oneight"}, 48},
		{"spelt overlap", args{"sevenine"}, 79},
		{"spelt overlap mix", args{"1sevenine"}, 19},
		{"mix", args{"7pqrstsixteen"}, 76},
		{"mix", args{"7pqrstsixteen2"}, 72},
		{"mix", args{"zoneight234"}, 14},
		{"mix", args{"4nineeightseven2"}, 42},
		{"mix", args{"xtwone3four"}, 24},
		{"mix", args{"two1nine"}, 29},
		{"mix", args{"abcone2threexyz"}, 13},
		{"mix", args{"szsvltgsc1onecccbfour3oneightfh"}, 18},
		{"mix", args{"xoneightnineoneltpvkzfvpqhxszqmv3hthtn"}, 13},
		{"mix", args{"xone1onetwo13four"}, 14},
		{"duplicate num mix", args{"rnkoneight1bvvhkd"}, 11},
		{"duplicate num mix", args{"rnk81oneightbvvhkd"}, 88},
		{"duplicate num mix", args{"rnkoneight17bvvhkd"}, 17},
		{"duplicate num mix", args{"rnk41one213bvvh1kd"}, 41},
		{"duplicate num mix", args{"xoneightnineoneltpvkzfvpqhxszqmv3hthtn"}, 13},
		{"duplicate num mix", args{"seven5rkmrc8six9oneights"}, 78},
		{"duplicate num mix", args{"three98oneightzn"}, 38},
		{"random test case", args{"three98oneightzn"}, 38},
		{"gtlbhbjgkrb5sixfivefivetwosix", args{"gtlbhbjgkrb5sixfivefivetwosix"}, 56},
		{"ninesixrgxccvrqscbskgzxh6cpvpxsqnb6", args{"ninesixrgxccvrqscbskgzxh6cpvpxsqnb6"}, 96},
		{"dxxzrlzkksfsffp4", args{"dxxzrlzkksfsffp4"}, 44},
		{"sbzvmddhnjtwollnjv33d2lbcscstqt", args{"sbzvmddhnjtwollnjv33d2lbcscstqt"}, 22},
		{"88xpnfpb", args{"88xpnfpb"}, 88},
		{"ninevct4cpdvqfxmspbz9xrvxfvbpzthreesfnncrqn", args{"ninevct4cpdvqfxmspbz9xrvxfvbpzthreesfnncrqn"}, 93},
		{"mqsxzsglbtmsbltrbzkbrt23", args{"mqsxzsglbtmsbltrbzkbrt23"}, 23},
		{"seven16ninefc", args{"seven16ninefc"}, 79},
		{"8jdddctvxs3", args{"8jdddctvxs3"}, 83},
		{"fivennhhdfpmrnpjhdm2sixkrsgdt", args{"fivennhhdfpmrnpjhdm2sixkrsgdt"}, 56},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcSum(tt.args.line); got != tt.want {
				t.Errorf("CalcSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

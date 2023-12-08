package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findStrongestHand(t *testing.T) {
	type args struct {
		h hand
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Joker Makes a five of a kind  - 5 J",
			args: args{h: hand{
				cards: [5]int{jack, jack, jack, jack, jack},
			}},
			want: fiveOfAKind,
		},
		{
			name: "Joker Makes a five of a kind  - 4 J",
			args: args{h: hand{
				cards: [5]int{queen, jack, jack, jack, jack},
			}},
			want: fiveOfAKind,
		},
		{
			name: "Joker Makes a five of a kind  - 3 J",
			args: args{h: hand{
				cards: [5]int{queen, queen, jack, jack, jack},
			}},
			want: fiveOfAKind,
		},
		{
			name: "Joker Makes a five of a kind  - 2 J",
			args: args{h: hand{
				cards: [5]int{queen, jack, queen, queen, jack},
			}},
			want: fiveOfAKind,
		},
		{
			name: "Joker Makes a five of a kind  - 1 J",
			args: args{h: hand{
				cards: [5]int{queen, queen, queen, queen, jack},
			}},
			want: fiveOfAKind,
		},
		{
			name: "No Joker Makes a five of a kind  - 0 J",
			args: args{h: hand{
				cards: [5]int{queen, queen, queen, queen, queen},
			}},
			want: fiveOfAKind,
		},
		{
			name: "No Joker Makes a four of a kind - 0 J",
			args: args{h: hand{
				cards: [5]int{queen, queen, queen, queen, two},
			}},
			want: fourOfAKind,
		},
		{
			name: "Joker Makes a four of a kind - 1 J",
			args: args{h: hand{
				cards: [5]int{queen, jack, queen, queen, two},
			}},
			want: fourOfAKind,
		},
		{
			name: "Joker Makes a four of a kind - 2 jokers",
			args: args{h: hand{
				cards: [5]int{queen, king, jack, king, jack},
			}},
			want: fourOfAKind,
		},
		{
			name: "Joker Makes a four of a kind - 3 jokers",
			args: args{h: hand{
				cards: [5]int{queen, jack, jack, king, jack},
			}},
			want: fourOfAKind,
		},
		{
			name: "Joker Makes a full house",
			args: args{h: hand{
				cards: [5]int{ace, jack, ace, queen, queen},
			}},
			want: fullHouse,
		},
		{
			name: "Joker Makes a three of a kind - 1 J",
			args: args{h: hand{
				cards: [5]int{queen, jack, five, queen, two},
			}},
			want: threeOfAKind,
		},
		{
			name: "Joker Makes a three of a kind - 2 J",
			args: args{h: hand{
				cards: [5]int{jack, jack, five, queen, two},
			}},
			want: threeOfAKind,
		},
		{
			name: "Joker Makes a three of a kind - 1 J",
			args: args{h: hand{
				cards: [5]int{queen, jack, five, queen, two},
			}},
			want: threeOfAKind,
		},
		{
			name: "No Joker Makes a three of a kind - 0 J",
			args: args{h: hand{
				cards: [5]int{queen, queen, five, queen, two},
			}},
			want: threeOfAKind,
		},
		{
			name: "No Joker Makes two pair - 0 J",
			args: args{h: hand{
				cards: [5]int{ace, ace, five, queen, five},
			}},
			want: twoPair,
		},
		{
			name: "Joker Makes a pair - 1 J",
			args: args{h: hand{
				cards: [5]int{ace, jack, five, queen, two},
			}},
			want: onePair,
		},
		{
			name: "No Joker Makes a pair - 0 J",
			args: args{h: hand{
				cards: [5]int{ace, ace, five, queen, two},
			}},
			want: onePair,
		},
		{
			name: "No joker gives high card",
			args: args{h: hand{
				cards: [5]int{ace, king, five, queen, two},
			}},
			want: highCard,
		},
		{
			name: "No joker gives full house",
			args: args{h: hand{
				cards: [5]int{ace, ace, ace, queen, queen},
			}},
			want: fullHouse,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findStrongestHandAccountingForJokers(tt.args.h), "findStrongestHandAccountingForJokers(%v)", tt.args.h)
		})
	}
}

func Test_orderHandsByStrengthWithJokers(t *testing.T) {
	type args struct {
		hands []hand
	}
	tests := []struct {
		name string
		args args
		want []hand
	}{
		{
			name: "competing type",
			args: args{hands: []hand{
				{cards: [5]int{three, jack, king, king, king}},
				{cards: [5]int{two, jack, ace, ace, ace}},
			}},
			want: []hand{
				{cards: [5]int{two, jack, ace, ace, ace}},
				{cards: [5]int{three, jack, king, king, king}},
			},
		},
		{ // for the purpose of breaking ties between two hands of the same type, J is always treated as J,
			// not the card it's pretending to be: JKKK2 is weaker than QQQQ2 because J is weaker than Q.
			name: "competing type 2",
			args: args{hands: []hand{
				{cards: [5]int{queen, queen, queen, queen, two}},
				{cards: [5]int{jack, king, king, king, two}},
			}},
			want: []hand{
				{cards: [5]int{jack, king, king, king, two}},
				{cards: [5]int{queen, queen, queen, queen, two}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orderHandsWithJokers(tt.args.hands); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("orderHandsByType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	hands := []hand{
		{cards: [5]int{three, two, ten, three, king}, bid: 765},   // one pair
		{cards: [5]int{ten, five, five, jack, five}, bid: 684},    // three of a kind
		{cards: [5]int{king, king, six, seven, seven}, bid: 28},   // two pair
		{cards: [5]int{king, ten, jack, jack, ten}, bid: 220},     // two pair
		{cards: [5]int{queen, queen, queen, jack, ace}, bid: 483}, // three of a kind
	}

	assert.Equal(t, 5905, calculateWinnings(orderHandsWithJokers(hands)))

}
func TestPart2_orderHandsWithJokers(t *testing.T) {
	hands := []hand{
		{cards: [5]int{three, two, ten, three, king}, bid: 765},   // one pair
		{cards: [5]int{ten, five, five, jack, five}, bid: 684},    // three of a kind
		{cards: [5]int{king, king, six, seven, seven}, bid: 28},   // two pair
		{cards: [5]int{king, ten, jack, jack, ten}, bid: 220},     // two pair
		{cards: [5]int{queen, queen, queen, jack, ace}, bid: 483}, // three of a kind
	}
	expectedHands := []hand{
		{cards: [5]int{three, two, ten, three, king}, bid: 765},   // one pair
		{cards: [5]int{king, king, six, seven, seven}, bid: 28},   // two pair
		{cards: [5]int{ten, five, five, jack, five}, bid: 684},    // three of a kind
		{cards: [5]int{queen, queen, queen, jack, ace}, bid: 483}, // three of a kind
		{cards: [5]int{king, ten, jack, jack, ten}, bid: 220},     // two pair
	}

	assert.Equal(t, expectedHands, orderHandsWithJokers(hands))

}

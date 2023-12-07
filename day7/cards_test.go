package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hand_calcStrength(t *testing.T) {
	type fields struct {
		hand hand
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "high card is strength 0",
			fields: fields{hand: hand{cards: [5]int{two, three, four, five, six}}},
			want:   0,
		},
		{
			name:   "one pair is strength 1",
			fields: fields{hand: hand{cards: [5]int{ace, two, three, ace, four}}},
			want:   1,
		},
		{
			name:   "two pair is strength 2",
			fields: fields{hand: hand{cards: [5]int{two, three, four, three, two}}},
			want:   2,
		},
		{
			name:   "three of a kind is strength 3",
			fields: fields{hand: hand{cards: [5]int{ten, ten, ten, nine, eight}}},
			want:   3,
		},
		{
			name:   "full house is strength 4",
			fields: fields{hand: hand{cards: [5]int{two, three, three, three, two}}},
			want:   4,
		},
		{
			name:   "four of a kind is strength 5",
			fields: fields{hand: hand{cards: [5]int{ace, ace, eight, ace, ace}}},
			want:   5,
		},
		{
			name:   "five of a kind is strength 6",
			fields: fields{hand: hand{cards: [5]int{ace, ace, ace, ace, ace}}},
			want:   6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hand{
				cards: tt.fields.hand.cards,
			}
			if got := h.calcStrength(); got != tt.want {
				t.Errorf("calcStrength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orderHandsByStrength(t *testing.T) {
	type args struct {
		hands []hand
	}
	tests := []struct {
		name string
		args args
		want []hand
	}{
		{
			name: "all different types",
			args: args{hands: []hand{
				{cards: [5]int{ten, ten, ten, nine, eight}},    // three of a kind
				{cards: [5]int{ace, two, three, ace, four}},    // one pair
				{cards: [5]int{ace, ace, eight, ace, ace}},     // four of a kind
				{cards: [5]int{two, three, four, three, two}},  // two pair
				{cards: [5]int{two, three, four, five, six}},   // high card
				{cards: [5]int{ace, ace, ace, ace, ace}},       // five of a kind
				{cards: [5]int{two, three, three, three, two}}, // full house
			}},
			want: []hand{
				{cards: [5]int{two, three, four, five, six}},   // high card
				{cards: [5]int{ace, two, three, ace, four}},    // one pair
				{cards: [5]int{two, three, four, three, two}},  // two pair
				{cards: [5]int{ten, ten, ten, nine, eight}},    // three of a kind
				{cards: [5]int{two, three, three, three, two}}, // full house
				{cards: [5]int{ace, ace, eight, ace, ace}},     // four of a kind
				{cards: [5]int{ace, ace, ace, ace, ace}},       // five of a kind
			},
		},
		{
			name: "competing type",
			args: args{hands: []hand{
				{cards: [5]int{two, three, four, five, six}},     // high card
				{cards: [5]int{three, three, three, three, two}}, // stronger four of a kind 33332
				{cards: [5]int{two, three, four, three, two}},    // two pair
				{cards: [5]int{ten, ten, ten, nine, eight}},      // three of a kind
				{cards: [5]int{two, three, three, three, two}},   // full house
				{cards: [5]int{two, ace, ace, ace, ace}},         // weaker four of a kind
				{cards: [5]int{ace, ace, ace, ace, ace}},         // five of a kind
			}},
			want: []hand{
				{cards: [5]int{two, three, four, five, six}},     // high card
				{cards: [5]int{two, three, four, three, two}},    // two pair
				{cards: [5]int{ten, ten, ten, nine, eight}},      // three of a kind
				{cards: [5]int{two, three, three, three, two}},   // full house
				{cards: [5]int{two, ace, ace, ace, ace}},         // weaker four of a kind
				{cards: [5]int{three, three, three, three, two}}, // stronger four of a kind 33332
				{cards: [5]int{ace, ace, ace, ace, ace}},         // five of a kind
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orderHands(tt.args.hands); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("orderHandsByType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateWinnings(t *testing.T) {
	type args struct {
		orderedHands []hand
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple case - note: already sorted",
			args: args{orderedHands: []hand{
				{cards: [5]int{three, two, ten, three, king}, bid: 765},   // one pair
				{cards: [5]int{king, ten, jack, jack, ten}, bid: 220},     // two pair
				{cards: [5]int{king, king, six, seven, seven}, bid: 28},   // two pair
				{cards: [5]int{ten, five, five, jack, five}, bid: 684},    // three of a kind
				{cards: [5]int{queen, queen, queen, jack, ace}, bid: 483}, // three of a kind
			}},
			want: 6440,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateWinnings(tt.args.orderedHands); got != tt.want {
				t.Errorf("calculateWinnings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	hands := []hand{
		{cards: [5]int{three, two, ten, three, king}, bid: 765},   // one pair
		{cards: [5]int{ten, five, five, jack, five}, bid: 684},    // three of a kind
		{cards: [5]int{king, king, six, seven, seven}, bid: 28},   // two pair
		{cards: [5]int{king, ten, jack, jack, ten}, bid: 220},     // two pair
		{cards: [5]int{queen, queen, queen, jack, ace}, bid: 483}, // three of a kind
	}

	assert.Equal(t, 6440, calculateWinnings(orderHands(hands)))

}

func Test_orderHandsByCard(t *testing.T) {
	type args struct {
		hands []hand
	}
	tests := []struct {
		name string
		args args
		want []hand
	}{
		{
			name: "simplest case",
			args: args{hands: []hand{
				{cards: [5]int{three, three, three, three, two}}, // stronger four of a kind
				{cards: [5]int{two, ace, ace, ace, ace}},         // weaker four of a kind
			},
			},
			want: []hand{
				{cards: [5]int{two, ace, ace, ace, ace}},         // weaker four of a kind
				{cards: [5]int{three, three, three, three, two}}, // stronger four of a kind
			},
		},
		{
			name: "three hands",
			args: args{hands: []hand{
				{cards: [5]int{four, four, four, four, two}}, // strongest four of a kind
				{cards: [5]int{three, ace, ace, ace, ace}},   // weaker four of a kind
				{cards: [5]int{two, king, king, king, king}}, // weakest four of a kind
			},
			},
			want: []hand{
				{cards: [5]int{two, king, king, king, king}}, // weakest four of a kind
				{cards: [5]int{three, ace, ace, ace, ace}},   // weaker four of a kind
				{cards: [5]int{four, four, four, four, two}}, // strongest four of a kind
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, orderHandsByCardForType(tt.args.hands), "orderHandsByCardForType(%v)", tt.args.hands)
		})
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part2(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	jack = 0 // Joker

	two = 1
	three = 2
	four = 3
	five = 4
	six = 5
	seven = 6
	eight = 7
	nine = 8
	ten = 9
	queen = 10
	king = 11
	ace = 12
	var hands []hand
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		rawHand, rawBet := lineSplit[0], lineSplit[1]
		bet, _ := strconv.Atoi(rawBet)
		cards := parseHand(rawHand)
		h := hand{
			cards: cards,
			bid:   bet,
		}
		hands = append(hands, h)
	}
	fmt.Println(calculateWinnings(orderHandsWithJokers(hands)))
}

func orderHandsWithJokers(hands []hand) []hand {
	orderHandsByTypeWithJokers(hands)

	// for each type, order those hands
	return orderHandsByCardForTypeWithJokers(hands)
}

func orderHandsByCardForTypeWithJokers(hands []hand) []hand {
	var highCardHands []hand
	var onePairHands []hand
	var twoPairHands []hand
	var threeOfAKindHands []hand
	var fullHouseHands []hand
	var fourOfAKindHands []hand
	var fiveOfAKindHands []hand

	for _, h := range hands {
		// set hands for each type
		switch findStrongestHandAccountingForJokers(h) {
		case highCard:
			highCardHands = append(highCardHands, h)
		case onePair:
			onePairHands = append(onePairHands, h)
		case twoPair:
			twoPairHands = append(twoPairHands, h)
		case threeOfAKind:
			threeOfAKindHands = append(threeOfAKindHands, h)
		case fullHouse:
			fullHouseHands = append(fullHouseHands, h)
		case fourOfAKind:
			fourOfAKindHands = append(fourOfAKindHands, h)
		case fiveOfAKind:
			fiveOfAKindHands = append(fiveOfAKindHands, h)
		}
	}

	sort.Sort(ByCardWithJoker(highCardHands))
	sort.Sort(ByCardWithJoker(onePairHands))
	sort.Sort(ByCardWithJoker(twoPairHands))
	sort.Sort(ByCardWithJoker(threeOfAKindHands))
	sort.Sort(ByCardWithJoker(fullHouseHands))
	sort.Sort(ByCardWithJoker(fourOfAKindHands))
	sort.Sort(ByCardWithJoker(fiveOfAKindHands))
	sortedCards := append(highCardHands, onePairHands...)
	sortedCards = append(sortedCards, twoPairHands...)
	sortedCards = append(sortedCards, threeOfAKindHands...)
	sortedCards = append(sortedCards, fullHouseHands...)
	sortedCards = append(sortedCards, fourOfAKindHands...)
	sortedCards = append(sortedCards, fiveOfAKindHands...)

	return sortedCards
}

type ByJokerType []hand

func (b ByJokerType) Len() int { return len(b) }

func (b ByJokerType) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByJokerType) Less(i, j int) bool {
	var iCards []int
	for _, c := range b[i].cards {
		iCards = append(iCards, c)
	}
	var jCards []int
	for _, c := range b[j].cards {
		jCards = append(jCards, c)
	}
	return findStrongestHandAccountingForJokers(b[i]) < findStrongestHandAccountingForJokers(b[j])
}

func orderHandsByTypeWithJokers(hands []hand) []hand {
	// first get hands in order of base strength
	sort.Sort(ByJokerType(hands))
	return hands
}

func calcStrengthPart2(cards []int) int {
	cardToOccurrence := map[int]int{}
	for _, card := range cards {
		cardToOccurrence[card]++
	}

	hasThreeOfAKind, hasOnePair := false, false
	for _, occurrences := range cardToOccurrence {
		if occurrences == 5 {
			return fiveOfAKind
		}
		if occurrences == 4 {
			return fourOfAKind
		}
		if occurrences == 3 {
			hasThreeOfAKind = true
			if hasOnePair {
				return fullHouse
			}
		}
		if occurrences == 2 && hasOnePair {
			return twoPair
		}
		if occurrences == 2 {
			hasOnePair = true
		}
	}

	if hasThreeOfAKind && hasOnePair {
		return fullHouse
	}

	if hasThreeOfAKind {
		return threeOfAKind
	}
	if hasOnePair {
		return onePair
	}

	return highCard
}

func findStrongestHandAccountingForJokers(h hand) int {
	// calc strength of hand without jokers
	var cardsSansJokers []int
	for _, c := range h.cards {
		if c == jack {
			continue
		}
		cardsSansJokers = append(cardsSansJokers, c)
	}

	baseStrength := calcStrengthPart2(cardsSansJokers)
	if len(cardsSansJokers) == 5 {
		return baseStrength
	}
	if len(cardsSansJokers) == 0 {
		// all 5 cards are jokers
		return fiveOfAKind
	}

	jokerCount := 5 - len(cardsSansJokers)

	if baseStrength == fourOfAKind && jokerCount == 1 {
		return fiveOfAKind
	}
	if baseStrength == threeOfAKind && jokerCount == 2 {
		return fiveOfAKind
	}
	if baseStrength == onePair && jokerCount == 3 {
		return fiveOfAKind
	}
	if baseStrength == highCard && jokerCount == 4 {
		return fiveOfAKind
	}

	if baseStrength == threeOfAKind && jokerCount == 1 {
		return fourOfAKind
	}
	if baseStrength == onePair && jokerCount == 2 {
		return fourOfAKind
	}
	if baseStrength == highCard && jokerCount == 3 {
		return fourOfAKind
	}

	if baseStrength == fourOfAKind && jokerCount == 0 {
		return fourOfAKind
	}

	if baseStrength == twoPair && jokerCount == 1 {
		return fullHouse
	}

	if baseStrength == highCard && jokerCount == 2 {
		return threeOfAKind
	}
	if baseStrength == onePair && jokerCount == 1 {
		return threeOfAKind
	}
	if baseStrength == threeOfAKind && jokerCount == 0 {
		return threeOfAKind
	}

	if baseStrength == twoPair && jokerCount == 0 {
		return twoPair
	}

	if baseStrength == onePair && jokerCount == 0 {
		return onePair
	}

	if baseStrength == highCard && jokerCount == 1 {
		return onePair
	}
	if baseStrength == highCard {
		return jokerCount + highCard
	}

	return highCard
}

type ByCardWithJoker []hand

func (b ByCardWithJoker) Len() int { return len(b) }

func (b ByCardWithJoker) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// for the purpose of breaking ties between two hands of the same type, J is always treated as J,
// not the card it's pretending to be: JKKK2 is weaker than QQQQ2 because J is weaker than Q.
func (b ByCardWithJoker) Less(i, j int) bool {
	for k := 0; k < 5; k++ {
		if b[i].cards[k] == b[j].cards[k] {
			continue
		} else {
			return b[i].cards[k] < b[j].cards[k]
		}
	}
	return false
}

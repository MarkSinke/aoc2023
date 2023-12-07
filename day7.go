package aoc2023

import (
	"slices"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

type Card int

func (x Card) String() string {
	return string(cardStrings[x])
}

const cardStrings = "23456789TJQKA"

func CardsFromString(str string) []Card {
	var cards []Card
	for _, r := range str {
		card := Card(strings.IndexRune(cardStrings, r))
		cards = append(cards, card)
	}
	return cards
}

type Hand struct {
	cards []Card
	rank  int
	bet   int
}

func getItem(m map[Card]int) (Card, int) {
	for c, i := range m {
		return c, i
	}
	return 0, 0
}

func makeRank(cards []Card) int {
	m := map[Card]int{}

	for _, card := range cards {
		entry, found := m[card]
		if !found {
			entry = 0
		}
		entry++
		m[card] = entry
	}

	var counts []int = maps.Values(m)
	// sort in highest-count first
	slices.SortFunc(counts, func(a int, b int) int { return b - a })

	switch counts[0] {
	case 5: // five of a kind
		return 10
	case 4: // four of a kind
		return 9
	case 3:
		if counts[1] == 2 {
			return 8 // full house
		} else {
			return 7 // three of a kind
		}
	case 2:
		if counts[1] == 2 {
			return 6 // two pairs
		} else {
			return 5 // one pair
		}
	case 1:
		return 4 // high card
	}

	// should not happen
	return 0
}

func SortHands(hands []Hand) {
	slices.SortFunc(hands, func(a Hand, b Hand) int {
		cmp := a.rank - b.rank
		if cmp != 0 {
			return cmp
		}

		for i := range a.cards {
			cmp := a.cards[i] - b.cards[i]
			if cmp != 0 {
				return int(cmp)
			}
		}
		return 0
	})
}

func ReadHands(path string) []Hand {
	lines := ReadFile(path)
	var hands []Hand
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		cards := CardsFromString(parts[0])
		rank := makeRank(cards)
		bet, _ := strconv.Atoi(parts[1])
		hands = append(hands, Hand{cards, rank, bet})
	}
	return hands
}

func TotalWinnings(hands []Hand) int {
	SortHands(hands)

	sum := 0
	for i, hand := range hands {
		sum += (i + 1) * hand.bet
	}
	return sum
}

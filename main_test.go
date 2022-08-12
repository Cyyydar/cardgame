package main

import (
	"testing"
)

func testEq(a, b []Card) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestDeckGenerator(t *testing.T) {
	deck := make([]Card, 52)
	deckGenerator(deck)
	trueDeck := []Card{{Suit: "heart", Value: 2}, {Suit: "heart", Value: 3}, {Suit: "heart", Value: 4}, {Suit: "heart", Value: 5}, {Suit: "heart", Value: 6}, {Suit: "heart", Value: 7}, {Suit: "heart", Value: 8}, {Suit: "heart", Value: 9}, {Suit: "heart", Value: 10}, {Suit: "heart", Value: 11}, {Suit: "heart", Value: 12}, {Suit: "heart", Value: 13}, {Suit: "heart", Value: 14}, {Suit: "diamond", Value: 2}, {Suit: "diamond", Value: 3}, {Suit: "diamond", Value: 4}, {Suit: "diamond", Value: 5}, {Suit: "diamond", Value: 6}, {Suit: "diamond", Value: 7}, {Suit: "diamond", Value: 8}, {Suit: "diamond", Value: 9}, {Suit: "diamond", Value: 10}, {Suit: "diamond", Value: 11}, {Suit: "diamond", Value: 12}, {Suit: "diamond", Value: 13}, {Suit: "diamond", Value: 14}, {Suit: "club", Value: 2}, {Suit: "club", Value: 3}, {Suit: "club", Value: 4}, {Suit: "club", Value: 5}, {Suit: "club", Value: 6}, {Suit: "club", Value: 7}, {Suit: "club", Value: 8}, {Suit: "club", Value: 9}, {Suit: "club", Value: 10}, {Suit: "club", Value: 11}, {Suit: "club", Value: 12}, {Suit: "club", Value: 13}, {Suit: "club", Value: 14}, {Suit: "spade", Value: 2}, {Suit: "spade", Value: 3}, {Suit: "spade", Value: 4}, {Suit: "spade", Value: 5}, {Suit: "spade", Value: 6}, {Suit: "spade", Value: 7}, {Suit: "spade", Value: 8}, {Suit: "spade", Value: 9}, {Suit: "spade", Value: 10}, {Suit: "spade", Value: 11}, {Suit: "spade", Value: 12}, {Suit: "spade", Value: 13}, {Suit: "spade", Value: 14}}
	if testEq(deck, trueDeck) {
		return
	}
	t.Error("Generation failed")
}

package main

import (
	"testing"
)

type testDataRemove struct {
	data   []Card
	want   []Card
	remove int
}

type testDataFill struct {
	data  []Card
	want  []Card
	trump Card
}

func testEq(a, b []Card) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[Card]int)
	for i := range a {
		m[a[i]]++
		m[b[i]]++
	}
	for _, o := range m {
		if o != 2 {
			return false
		}
	}
	return true
}

func Find(a []Card, x Card) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

func TestDeckGenerator(t *testing.T) {
	Deck = make([]Card, 52)
	DeckGenerator()
	trueDeck := []Card{{Suit: "heart", Value: 2}, {Suit: "heart", Value: 3}, {Suit: "heart", Value: 4}, {Suit: "heart", Value: 5}, {Suit: "heart", Value: 6}, {Suit: "heart", Value: 7}, {Suit: "heart", Value: 8}, {Suit: "heart", Value: 9}, {Suit: "heart", Value: 10}, {Suit: "heart", Value: 11}, {Suit: "heart", Value: 12}, {Suit: "heart", Value: 13}, {Suit: "heart", Value: 14}, {Suit: "diamond", Value: 2}, {Suit: "diamond", Value: 3}, {Suit: "diamond", Value: 4}, {Suit: "diamond", Value: 5}, {Suit: "diamond", Value: 6}, {Suit: "diamond", Value: 7}, {Suit: "diamond", Value: 8}, {Suit: "diamond", Value: 9}, {Suit: "diamond", Value: 10}, {Suit: "diamond", Value: 11}, {Suit: "diamond", Value: 12}, {Suit: "diamond", Value: 13}, {Suit: "diamond", Value: 14}, {Suit: "club", Value: 2}, {Suit: "club", Value: 3}, {Suit: "club", Value: 4}, {Suit: "club", Value: 5}, {Suit: "club", Value: 6}, {Suit: "club", Value: 7}, {Suit: "club", Value: 8}, {Suit: "club", Value: 9}, {Suit: "club", Value: 10}, {Suit: "club", Value: 11}, {Suit: "club", Value: 12}, {Suit: "club", Value: 13}, {Suit: "club", Value: 14}, {Suit: "spade", Value: 2}, {Suit: "spade", Value: 3}, {Suit: "spade", Value: 4}, {Suit: "spade", Value: 5}, {Suit: "spade", Value: 6}, {Suit: "spade", Value: 7}, {Suit: "spade", Value: 8}, {Suit: "spade", Value: 9}, {Suit: "spade", Value: 10}, {Suit: "spade", Value: 11}, {Suit: "spade", Value: 12}, {Suit: "spade", Value: 13}, {Suit: "spade", Value: 14}}
	trueDeck = Remove(trueDeck, Find(trueDeck, trumpcard))
	if testEq(Deck, trueDeck) {
		return
	}
	t.Error("Generation failed\n", Deck, "\n", trueDeck, "\n", trumpcard)
}

func TestRemoveDeck(t *testing.T) {
	tests := []testDataRemove{
		/*1*/ {[]Card{{Suit: "heart", Value: 2}, {Suit: "heart", Value: 3}, {Suit: "heart", Value: 4}, {Suit: "heart", Value: 5}, {Suit: "heart", Value: 6}},
			[]Card{{Suit: "heart", Value: 3}, {Suit: "heart", Value: 4}, {Suit: "heart", Value: 5}, {Suit: "heart", Value: 6}},
			0},
		/*2*/ {[]Card{{Suit: "heart", Value: 2}, {Suit: "heart", Value: 3}, {Suit: "heart", Value: 4}, {Suit: "heart", Value: 5}, {Suit: "heart", Value: 6}},
			[]Card{{Suit: "heart", Value: 2}, {Suit: "heart", Value: 3}, {Suit: "heart", Value: 5}, {Suit: "heart", Value: 6}},
			2},
		/*3*/ {[]Card{{Suit: "heart", Value: 2}, {Suit: "heart", Value: 3}, {Suit: "heart", Value: 4}, {Suit: "heart", Value: 5}, {Suit: "heart", Value: 6}},
			[]Card{{Suit: "heart", Value: 2}, {Suit: "heart", Value: 3}, {Suit: "heart", Value: 4}, {Suit: "heart", Value: 5}},
			4},
		/*4*/ {[]Card{{Suit: "heart", Value: 2}, {Suit: "heart", Value: 3}, {Suit: "heart", Value: 4}, {Suit: "heart", Value: 5}, {Suit: "heart", Value: 6}, {Suit: "heart", Value: 7}, {Suit: "heart", Value: 8}, {Suit: "heart", Value: 9}, {Suit: "heart", Value: 10}, {Suit: "heart", Value: 11}, {Suit: "heart", Value: 12}, {Suit: "heart", Value: 13}, {Suit: "heart", Value: 14}, {Suit: "diamond", Value: 2}, {Suit: "diamond", Value: 3}, {Suit: "diamond", Value: 4}, {Suit: "diamond", Value: 5}, {Suit: "diamond", Value: 6}, {Suit: "diamond", Value: 7}, {Suit: "diamond", Value: 8}, {Suit: "diamond", Value: 9}, {Suit: "diamond", Value: 10}, {Suit: "diamond", Value: 11}, {Suit: "diamond", Value: 12}, {Suit: "diamond", Value: 13}, {Suit: "diamond", Value: 14}, {Suit: "club", Value: 2}, {Suit: "club", Value: 3}, {Suit: "club", Value: 4}, {Suit: "club", Value: 5}, {Suit: "club", Value: 6}, {Suit: "club", Value: 7}, {Suit: "club", Value: 8}, {Suit: "club", Value: 9}, {Suit: "club", Value: 10}, {Suit: "club", Value: 11}, {Suit: "club", Value: 12}, {Suit: "club", Value: 13}, {Suit: "club", Value: 14}, {Suit: "spade", Value: 2}, {Suit: "spade", Value: 3}, {Suit: "spade", Value: 4}, {Suit: "spade", Value: 5}, {Suit: "spade", Value: 6}, {Suit: "spade", Value: 7}, {Suit: "spade", Value: 8}, {Suit: "spade", Value: 9}, {Suit: "spade", Value: 10}, {Suit: "spade", Value: 11}, {Suit: "spade", Value: 12}, {Suit: "spade", Value: 13}, {Suit: "spade", Value: 14}},
			[]Card{{Suit: "heart", Value: 3}, {Suit: "heart", Value: 4}, {Suit: "heart", Value: 5}, {Suit: "heart", Value: 6}, {Suit: "heart", Value: 7}, {Suit: "heart", Value: 8}, {Suit: "heart", Value: 9}, {Suit: "heart", Value: 10}, {Suit: "heart", Value: 11}, {Suit: "heart", Value: 12}, {Suit: "heart", Value: 13}, {Suit: "heart", Value: 14}, {Suit: "diamond", Value: 2}, {Suit: "diamond", Value: 3}, {Suit: "diamond", Value: 4}, {Suit: "diamond", Value: 5}, {Suit: "diamond", Value: 6}, {Suit: "diamond", Value: 7}, {Suit: "diamond", Value: 8}, {Suit: "diamond", Value: 9}, {Suit: "diamond", Value: 10}, {Suit: "diamond", Value: 11}, {Suit: "diamond", Value: 12}, {Suit: "diamond", Value: 13}, {Suit: "diamond", Value: 14}, {Suit: "club", Value: 2}, {Suit: "club", Value: 3}, {Suit: "club", Value: 4}, {Suit: "club", Value: 5}, {Suit: "club", Value: 6}, {Suit: "club", Value: 7}, {Suit: "club", Value: 8}, {Suit: "club", Value: 9}, {Suit: "club", Value: 10}, {Suit: "club", Value: 11}, {Suit: "club", Value: 12}, {Suit: "club", Value: 13}, {Suit: "club", Value: 14}, {Suit: "spade", Value: 2}, {Suit: "spade", Value: 3}, {Suit: "spade", Value: 4}, {Suit: "spade", Value: 5}, {Suit: "spade", Value: 6}, {Suit: "spade", Value: 7}, {Suit: "spade", Value: 8}, {Suit: "spade", Value: 9}, {Suit: "spade", Value: 10}, {Suit: "spade", Value: 11}, {Suit: "spade", Value: 12}, {Suit: "spade", Value: 13}, {Suit: "spade", Value: 14}},
			0},
	}
	for _, test := range tests {
		got := Remove(test.data, test.remove)
		if !testEq(got, test.want) {
			t.Errorf("TestRemoveDeck() =\n \"%v\" \n, want \n \"%v\" ", got, test.want)
		}
	}
}

func TestFill(t *testing.T) {
	//Deck = []Card{{Suit: "heart", Value: 2}, {Suit: "heart", Value: 3}, {Suit: "heart", Value: 4}, {Suit: "heart", Value: 5}, {Suit: "heart", Value: 6}, {Suit: "heart", Value: 7}}
	//trueHand = []Card{{Suit: "heart", Value: 2}, {Suit: "heart", Value: 3}, {Suit: "heart", Value: 4}, {Suit: "heart", Value: 5}, {Suit: "heart", Value: 6}, {Suit: "heart", Value: 7}}
	// hand := Hand{}
	// hand.Fill()
	// if len(Deck) != 0 || len(hand.Hand) != 6 {
	// 	t.Error("Deck: ", Deck, "hand: ", hand)
	// }
	tests := []testDataFill{
		{[]Card{{Suit: "heart", Value: 2}, {Suit: "heart", Value: 3}, {Suit: "heart", Value: 4}, {Suit: "heart", Value: 5}, {Suit: "heart", Value: 6}, {Suit: "heart", Value: 7}},
			[]Card{{Suit: "heart", Value: 2}, {Suit: "heart", Value: 3}, {Suit: "heart", Value: 4}, {Suit: "heart", Value: 5}, {Suit: "heart", Value: 6}, {Suit: "heart", Value: 7}},
			Card{Suit: "heart", Value: 14}},
		{[]Card{},
			[]Card{{Suit: "heart", Value: 14}},
			Card{Suit: "heart", Value: 14}},
	}
	for _, test := range tests {
		Deck = test.data
		trumpcard = test.trump
		hand := Hand{}
		hand.Fill()
		if !testEq(hand.Hand, test.want) {
			t.Errorf("TestRemoveDeck() =\n \"%v\" \n, want \n \"%v\" ", hand.Hand, test.want)
		}
	}
}

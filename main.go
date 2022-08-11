package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func RandomNumber(c int) int {
	seed := time.Now().Unix()
	rand.Seed(seed)
	number := rand.Intn(c)
	return number
}

func Remove(d []Card, c int) []Card {
	if c > len(d) {
		log.Fatal("Мать жива? Ты нафига удаляешь не ту карту")
	}
	d = append(d[:c], d[c+1:]...)
	return d
}

func deckGenerator(d []Card) {
	suits := []string{"heart", "diamond", "club", "spade"}
	number := 0
	for _, suit := range suits {
		for index, value := number, 2; index < number+13; index++ {
			d[index] = Card{suit, value}
			value++
		}
		number += 13
	}
}

type Card struct {
	Suit  string
	Value int
}

type Hand struct {
	hand []Card
}

func (h *Hand) get(d []Card) {
	if 6-len(h.hand) < 0 || len(d) == 0 {
		return
	}
	for i := len(h.hand); i < 6; i++ {
		h.hand = append(h.hand, d[RandomNumber(len(d))])
		d = Remove(d, i)
	}
}

func (h *Hand) attack(anotherHand *Hand) {
	var cardnum int
	deflen := len(anotherHand.hand)
	for i := 0; i < deflen; i++ {
		fmt.Println("Attack: Choice card")
		fmt.Println(h.hand)

		fmt.Scanf("\n%d", &cardnum)
		h.hand = Remove(h.hand, cardnum)

		fmt.Println("Deffence: Choice card")
		fmt.Println(anotherHand.hand)

		fmt.Scanf("\n%d", &cardnum)
		anotherHand.hand = Remove(anotherHand.hand, cardnum)

	}
}

func main() {
	deck := make([]Card, 52)

	deckGenerator(deck)

	myHand := Hand{}
	myHand.get(deck)

	yourHand := Hand{}
	yourHand.get(deck)

	myHand.attack(&yourHand)
}

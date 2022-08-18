package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var Deck []Card
var attackTable []Card
var deffenceTable []Card

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

func DeckGenerator(d []Card) {
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
	Hand []Card
}

// заполняет руку 6 картами
func (h *Hand) Fill() {
	if 6-len(h.Hand) < 0 || len(Deck) == 0 {
		return
	}
	for i := len(h.Hand); i < 6; i++ {
		r := RandomNumber(len(Deck))
		h.Hand = append(h.Hand, Deck[r])
		Deck = Remove(Deck, r)
	}
}

// Возвращаем мапу номеров карт настоле для подкидывания
func Table(m map[int]int) {
	for _, card := range attackTable {
		m[card.Value]++
	}
	for _, card := range deffenceTable {
		m[card.Value]++
	}
}

// Запрашивает номер карты, удаляет карту и возвращает ее
func (h *Hand) pick() Card {
	cardnum := 1
	var choice string
	table := make(map[int]int)
	Table(table)

	fmt.Println(h.Hand) // печатаем руку
	for {
		fmt.Scanf("%s\n", &choice) //ожидаем номер карты
		if choice == "-" {
			return Card{"", 0}
		}

		cardnum, _ = strconv.Atoi(choice)
		cardnum--
		// if err != nil {
		// 	log.Fatal(err)
		// }
		if cardnum >= 0 && cardnum < len(h.Hand) { // если данные введены корректно
			fmt.Println(table)
			_, ok := table[h.Hand[cardnum].Value]
			if len(attackTable) <= len(deffenceTable) && (ok || len(table) == 0) { // атака
				fmt.Println("Attack")
				fmt.Println("\033[31m", h.Hand[cardnum], "\033[0m") // печатаем карту
				card := h.Hand[cardnum]                             // запоминаем карту
				h.Hand = Remove(h.Hand, cardnum)                    //удаляем карту из руки

				return card
			} else if len(attackTable) > len(deffenceTable) && h.Hand[cardnum].Value > attackTable[len(attackTable)-1].Value && h.Hand[cardnum].Suit == attackTable[len(attackTable)-1].Suit { //защита
				fmt.Println("\033[31m", h.Hand[cardnum], "\033[0m") // печатаем карту
				card := h.Hand[cardnum]                             // запоминаем карту
				h.Hand = Remove(h.Hand, cardnum)                    //удаляем карту из руки

				return card
			}
		}
		fmt.Println("Please try again")
	}

}

func (h *Hand) Attack(anotherHand *Hand) bool {
	attackTable = nil
	deffenceTable = nil
	deflen := len(anotherHand.Hand)
	skip := false

	for i := 0; i < deflen; i++ {
		fmt.Println("\033[32mAttack: Choice card", 1, "-", len(h.Hand), "\033[0m")
		attackTable = append(attackTable, h.pick())
		if attackTable[len(attackTable)-1].Value == 0 {
			fmt.Println("End of attack")
			break
		}
		fmt.Println(attackTable)   //печатаем столы
		fmt.Println(deffenceTable) //печатаем столы

		fmt.Println("\033[32mDeffence: Choice card", 1, "-", len(anotherHand.Hand), "\033[0m")
		deffenceTable = append(deffenceTable, anotherHand.pick())
		if deffenceTable[len(deffenceTable)-1].Value == 0 {
			fmt.Println("You lose")
			anotherHand.Hand = append(anotherHand.Hand, deffenceTable[:len(deffenceTable)-1]...)
			anotherHand.Hand = append(anotherHand.Hand, attackTable...)
			skip = true
			break
		}
	}
	// добираем карты
	h.Fill()
	anotherHand.Fill()
	return skip
}

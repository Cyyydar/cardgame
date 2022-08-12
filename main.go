package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
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

// заполняет руку 6 картами
func (h *Hand) fill() {
	if 6-len(h.hand) < 0 || len(Deck) == 0 {
		return
	}
	for i := len(h.hand); i < 6; i++ {
		h.hand = append(h.hand, Deck[RandomNumber(len(Deck))])
		Deck = Remove(Deck, i)
	}
}

// Запрашивает номер карты, удаляет карту и возвращает ее
func (h *Hand) pick() Card {
	cardnum := 1
	var choice string
	fmt.Println(h.hand) // печатаем руку
	for {
		fmt.Scanf("%s\n", &choice) //ожидаем номер карты
		if choice == "-" {
			return Card{"", 0}
		}

		cardnum, _ = strconv.Atoi(choice)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		if cardnum > 0 && cardnum <= len(h.hand) {
			fmt.Println("\033[31m", h.hand[cardnum-1], "\033[0m") // печатаем карту
			card := h.hand[cardnum-1]                             // запоминаем карту
			h.hand = Remove(h.hand, cardnum-1)                    //удаляем карту из руки

			return card
		}
		fmt.Println("Please try again")
	}

}

func (h *Hand) attack(anotherHand *Hand) {
	var attackTable []Card
	var deffenceTable []Card
	deflen := len(anotherHand.hand)
	for i := 0; i < deflen; i++ {
		fmt.Println("\033[32mAttack: Choice card", 1, "-", len(h.hand), "\033[0m")
		attackTable = append(attackTable, h.pick())
		if attackTable[len(attackTable)-1].Value == 0 {
			fmt.Println("End of attack")
			break
		}
		fmt.Println(attackTable)   //печатаем столы
		fmt.Println(deffenceTable) //печатаем столы

		fmt.Println("\033[32mDeffence: Choice card", 1, "-", len(anotherHand.hand), "\033[0m")
		deffenceTable = append(deffenceTable, anotherHand.pick())
		if deffenceTable[len(deffenceTable)-1].Value == 0 {
			fmt.Println("You lose")
			anotherHand.hand = append(anotherHand.hand, deffenceTable[:len(deffenceTable)-1]...)
			anotherHand.hand = append(anotherHand.hand, attackTable...)
			break
		}
	}

	//h. fill(Deck)
	//anotherHand. fill(Deck)
}

var Deck []Card

func main() {
	Deck = make([]Card, 52)
	deckGenerator(Deck)
	fmt.Println(Deck, len(Deck))

	countPlayers := 0 // спрашиваем сколько игроков будет
	for {
		fmt.Println("Enter count players")
		var temp string
		fmt.Scanf("%s\n", &temp)
		countPlayers, _ = strconv.Atoi(temp)
		if countPlayers > 1 && countPlayers < 10 {
			break
		}
		fmt.Println("Try again")
	}
	players := make([]Hand, countPlayers) // создаем слайс игроков

	for i := range players {
		players[i].fill()
	}
	fmt.Println(Deck, len(Deck))

	// атака по очереди
	for i := 0; ; i++ {
		fmt.Println("Turn player", i)
		if i == len(players)-1 {
			players[i].attack(&players[0])
			i = -1
		} else {
			players[i].attack(&players[i+1])
		}
	}

	// myHand := Hand{}
	// myHand. fill(Deck)

	// yourHand := Hand{}
	// yourHand. fill(Deck)

	//myHand.attack(&yourHand)

}

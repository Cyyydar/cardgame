package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// --------------------------------
var Deck []Card
var attackTable []Card
var defenseTable []Card
var trumpcard Card
var players []Hand

func RandomNumber(c int) int {
	seed := time.Now().Unix()
	rand.Seed(seed)
	number := rand.Intn(c)
	return number
}

func Remove(c []Card, i int) []Card {
	if i > len(c) {
		log.Fatal("Мать жива? Ты нафига удаляешь не ту карту")
	}
	c = append(c[:i], c[i+1:]...)
	return c
}

func DeckGenerator() {
	suits := []string{"heart", "diamond", "club", "spade"}
	number := 0
	for _, suit := range suits {
		for index, value := number, 2; index < number+13; index++ {
			Deck[index] = Card{suit, value}
			value++
		}
		number += 13
	}
	r := RandomNumber(len(Deck))
	trumpcard = Deck[r]
	Deck = Remove(Deck, r)
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
	if (6-len(h.Hand) < 0 || len(Deck) == 0) && trumpcard != (Card{"", 0}) {
		h.Hand = append(h.Hand, trumpcard)
		trumpcard = Card{"", 0}
		return
	} else if trumpcard == (Card{"", 0}) {
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
	for _, card := range defenseTable {
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
			_, ok := table[h.Hand[cardnum].Value]
			if len(attackTable) <= len(defenseTable) && (ok || len(table) == 0) { // атака
				fmt.Println("Attack")
				fmt.Println("\033[31m", h.Hand[cardnum], "\033[0m") // печатаем карту
				card := h.Hand[cardnum]                             // запоминаем карту
				h.Hand = Remove(h.Hand, cardnum)                    //удаляем карту из руки

				return card
			} else if len(attackTable) > len(defenseTable) && h.Hand[cardnum].Value > attackTable[len(attackTable)-1].Value && h.Hand[cardnum].Suit == attackTable[len(attackTable)-1].Suit || h.Hand[cardnum].Suit == trumpcard.Suit && attackTable[len(attackTable)-1].Suit != trumpcard.Suit { //защита
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
	defenseTable = nil
	deflen := len(anotherHand.Hand)
	skip := false

	for i := 0; i < deflen; i++ {
		fmt.Println("\033[32mAttack: Choice card", 1, "-", len(h.Hand), "\033[0m")
		attackTable = append(attackTable, h.pick())
		if attackTable[len(attackTable)-1].Value == 0 {
			fmt.Println("End of attack")
			break
		}
		fmt.Println(attackTable)  //печатаем столы
		fmt.Println(defenseTable) //печатаем столы

		fmt.Println("\033[32mDefense: Choice card", 1, "-", len(anotherHand.Hand), "\033[0m")
		defenseTable = append(defenseTable, anotherHand.pick())
		if defenseTable[len(defenseTable)-1].Value == 0 {
			fmt.Println("You lose")
			anotherHand.Hand = append(anotherHand.Hand, defenseTable[:len(defenseTable)-1]...)
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

//--------------------------------

func skip(b bool, i *int) {
	if b {
		*i++
	}
}

func start(countPlayers int) {
	Deck = make([]Card, 52)
	DeckGenerator()
	//countPlayers := 0 // спрашиваем сколько игроков будет
	// for {
	// 	fmt.Println("Enter the number of players")
	// 	var temp string
	// 	fmt.Scanf("%s\n", &temp)
	// 	countPlayers, _ = strconv.Atoi(temp)
	// 	if countPlayers > 1 && countPlayers < 10 {
	// 		break
	// 	}
	// 	fmt.Println("Try again")
	// }
	players = make([]Hand, countPlayers) // создаем слайс игроков

	for i := range players {
		players[i].Fill()
	}
}

func main() {
	web()
	//start()
	fmt.Println(trumpcard)
	// атака по кругу
	for i := 0; ; i++ {
		fmt.Println("Turn player", i%len(players), "under", (i+1)%len(players))
		sp := players[i%len(players)].Attack(&players[(i+1)%len(players)])
		skip(sp, &i)
	}

}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func startHandler(writer http.ResponseWriter, request *http.Request) {
	html, _ := template.ParseFiles("start.html")
	html.Execute(writer, nil)

}

func gameHandler(writer http.ResponseWriter, request *http.Request) {
	type Data struct {
		Trumpcard Card
		AttackTable []Card
		DefenseTable []Card
		Hand []Card
	}
	gameData := Data{}
	gameData.Hand = []Card{{Suit: "heart", Value: 2}, {Suit: "heart", Value: 3}, {Suit: "heart", Value: 4}, {Suit: "heart", Value: 5}, {Suit: "heart", Value: 6}, {Suit: "heart", Value: 7}}
	gameData.AttackTable = gameData.Hand
	gameData.DefenseTable = gameData.Hand
	gameData.Trumpcard = Card{"heart",2}
	html, _ := template.ParseFiles("main.html")
	html.Execute(writer, gameData)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	temp := request.FormValue("countPlayers")
	//_, err := writer.Write([]byte(temp))
	//check(err)

	countPlayers := 0
	countPlayers, _ = strconv.Atoi(temp)
	if countPlayers > 0 && countPlayers <= 6{
		start(countPlayers)
		http.Redirect(writer,request, "/game", http.StatusFound)
	} else {
		http.Redirect(writer,request, "/game/start", http.StatusFound)
	}
}

func turnHandler(writer http.ResponseWriter, request *http.Request) {
	temp := request.FormValue("cardNumber")
}

func web() {
	http.HandleFunc("/game/start", startHandler)
	http.HandleFunc("/game", gameHandler)
	http.HandleFunc("/game/start/create", createHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

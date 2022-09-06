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
var playerTurn = 1

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
	for i := len(h.Hand); i < 6 && len(Deck) > 0; i++ {
		r := RandomNumber(len(Deck))
		h.Hand = append(h.Hand, Deck[r])
		Deck = Remove(Deck, r)
	}
	if 6-len(h.Hand) < 0 && len(Deck) == 0 && trumpcard != (Card{"", 0}) {
		h.Hand = append(h.Hand, trumpcard)
		trumpcard = Card{"", 0}
		return
	} else if trumpcard == (Card{"", 0}) {
		return
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
			} else if len(attackTable) > len(defenseTable) && h.Hand[cardnum].Value > attackTable[len(attackTable)-1].Value && h.Hand[cardnum].Suit == attackTable[len(attackTable)-1].Suit || attackTable[len(attackTable)-1].Suit != trumpcard.Suit && h.Hand[cardnum].Suit == trumpcard.Suit { //защита
				fmt.Println("\033[31m", h.Hand[cardnum], "\033[0m") // печатаем карту
				card := h.Hand[cardnum]                             // запоминаем карту
				h.Hand = Remove(h.Hand, cardnum)                    //удаляем карту из руки

				return card
			}
		}
		fmt.Println("Please try again")
	}

}

// func Attack() {
// 	h = &players[playerTurn%len(players)]
// 	anotherHand = &players[(playerTurn+1)%len(players)]
// 	//sp := players[playerTurn%len(players)].Attack(&players[(playerTurn+1)%len(players)])

// 	deflen := len(anotherHand.Hand)
// 	skip := false

// 	for i := 0; i < deflen; i++ {
// 		fmt.Println("\033[32mAttack: Choice card", 1, "-", len(h.Hand), "\033[0m")
// 		attackTable = append(attackTable, h.pick())
// 		if attackTable[len(attackTable)-1].Value == 0 {
// 			fmt.Println("End of attack")
// 			break
// 		}
// 		fmt.Println(attackTable)  //печатаем столы
// 		fmt.Println(defenseTable) //печатаем столы

// 		fmt.Println("\033[32mDefense: Choice card", 1, "-", len(anotherHand.Hand), "\033[0m")
// 		defenseTable = append(defenseTable, anotherHand.pick())
// 		if defenseTable[len(defenseTable)-1].Value == 0 {
// 			fmt.Println("You lose")
// 			anotherHand.Hand = append(anotherHand.Hand, defenseTable[:len(defenseTable)-1]...)
// 			anotherHand.Hand = append(anotherHand.Hand, attackTable...)

// 			playerTurn++
// 			attackTable = nil
// 			defenseTable = nil
// 			break
// 		}
// 	}
// 	// добираем карты
// 	h.Fill()
// 	anotherHand.Fill()
// }

func Turn(choice string) {
	// var choice string
	cardnum, _ := strconv.Atoi(choice)
	table := make(map[int]int)
	Table(table)
	fmt.Println(playerTurn)
	_, ok := table[players[playerTurn%len(players)].Hand[cardnum].Value]
	fmt.Println(playerTurn)

	if choice == "-" && len(attackTable) <= len(defenseTable) {
		endTurn()
		return
	} else if choice == "-" && len(attackTable) > len(defenseTable) {
		if len(defenseTable) != 0 {
			players[playerTurn%len(players)].Hand = append(players[playerTurn%len(players)].Hand, defenseTable[:len(defenseTable)-1]...)
		}
		players[playerTurn%len(players)].Hand = append(players[playerTurn%len(players)].Hand, attackTable...)
		endTurn()
		return
	}
	/* Нужно переработать систему с ходами:
	счетчик ходов +1 будет указывать на человека, который защищается
	или ввести отдельный счетчик защиты*/
	if len(attackTable) <= len(defenseTable) && (ok || len(table) == 0) { // атака
		fmt.Println("Attack")
		fmt.Println("\033[31m", players[playerTurn%len(players)].Hand[cardnum], "\033[0m") // печатаем карту
		attackTable = append(attackTable, players[playerTurn%len(players)].Hand[cardnum])
		players[playerTurn%len(players)].Hand = Remove(players[playerTurn%len(players)].Hand, cardnum)
		playerTurn++ //удаляем карту из руки
	} else if len(attackTable) > len(defenseTable) && players[playerTurn%len(players)].Hand[cardnum].Value > attackTable[len(attackTable)-1].Value && players[playerTurn%len(players)].Hand[cardnum].Suit == attackTable[len(attackTable)-1].Suit || players[playerTurn%len(players)].Hand[cardnum].Suit == trumpcard.Suit && attackTable[len(attackTable)-1].Suit != trumpcard.Suit { //защита
		fmt.Println("\033[31m", players[playerTurn%len(players)].Hand[cardnum], "\033[0m")             // печатаем карту
		defenseTable = append(defenseTable, players[playerTurn%len(players)].Hand[cardnum])            // запоминаем карту
		players[playerTurn%len(players)].Hand = Remove(players[playerTurn%len(players)].Hand, cardnum) //удаляем карту из руки
		playerTurn--
	}
}

func endTurn() {
	players[playerTurn%len(players)].Fill()
	players[(playerTurn+1)%len(players)].Fill()
	attackTable = nil
	defenseTable = nil
	playerTurn++
}

//--------------------------------

func skip(b bool) {
	if b {
		playerTurn++
	}
}

func start(countPlayers int) {
	attackTable = nil
	defenseTable = nil
	playerTurn = 0

	Deck = make([]Card, 52)
	DeckGenerator()
	players = make([]Hand, countPlayers) // создаем слайс игроков

	for i := range players {
		players[i].Fill()
	}
}

func main() {
	web()
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
	// if len(players) == 0 {
	// 	http.Redirect(writer,request, "/game/start", http.StatusFound)
	// }
	temp := string(request.URL.String()[len(request.URL.String())-1])
	currentPlayer, _ := strconv.Atoi(temp)
	type Data struct {
		Trumpcard     Card
		Turn          int
		AttackTable   []Card
		DefenseTable  []Card
		Hand          []Card
		CountCards    int
		CurrentPlayer int
	}
	gameData := Data{
		trumpcard,
		playerTurn % len(players),
		attackTable,
		defenseTable,
		players[currentPlayer].Hand,
		len(Deck),
		currentPlayer,
	}
	html, _ := template.ParseFiles("main.html")
	html.Execute(writer, gameData)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	temp := request.FormValue("countPlayers")
	//_, err := writer.Write([]byte(temp))
	//check(err)

	countPlayers := 0
	countPlayers, _ = strconv.Atoi(temp)
	if countPlayers > 0 && countPlayers <= 6 {
		start(countPlayers)
		for i := 0; i < countPlayers; i++ {
			http.HandleFunc("/game/"+fmt.Sprint(i), gameHandler)
		}
		http.Redirect(writer, request, "/game/start/room", http.StatusFound)
	} else {
		http.Redirect(writer, request, "/game/start", http.StatusFound)
	}
}

func roomHandler(writer http.ResponseWriter, request *http.Request) {
	type Data struct {
		Players []Hand
	}
	gameData := Data{
		Players: players,
	}
	html, _ := template.ParseFiles("room.html")
	html.Execute(writer, gameData)
}

func turnHandler(writer http.ResponseWriter, request *http.Request) {
	temp := request.FormValue("cardNumber")

	//fmt.Println("Turn player", playerTurn%len(players), "under", (playerTurn+1)%len(players))
	//sp := players[playerTurn%len(players)].Attack(&players[(playerTurn+1)%len(players)])
	// skip(sp)
	fmt.Println(playerTurn)
	Turn(temp)
	http.Redirect(writer, request, "/game", http.StatusFound)
}

func web() {
	http.HandleFunc("/game/start", startHandler)
	http.HandleFunc("/game/start/create", createHandler)
	http.HandleFunc("/game/start/room", roomHandler)
	http.HandleFunc("/game/turn", turnHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

package main

import (
	"fmt"
	"strconv"

	"github.com/Cyyydar/cards"
)

func skip(b bool, i *int) {
	if b {
		*i++
	}
}

func main() {
	cards.Deck = make([]cards.Card, 52)
	cards.DeckGenerator(cards.Deck)

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
	players := make([]cards.Hand, countPlayers) // создаем слайс игроков

	for i := range players {
		players[i].Fill()
	}

	// атака по кругу
	for i := 0; ; i++ {

		fmt.Println("Turn player", i)
		if i >= len(players)-1 {
			sp := players[i].Attack(&players[0])
			i = -1
			skip(sp, &i)

		} else {
			sp := players[i].Attack(&players[i+1])
			skip(sp, &i)
		}
	}

}

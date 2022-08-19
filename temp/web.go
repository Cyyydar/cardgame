package main

import (
	"html/template"
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, equest *http.Request) {
	test := []Card{{Suit: "heart", Value: 2}, {Suit: "heart", Value: 3}, {Suit: "heart", Value: 4}, {Suit: "heart", Value: 5}, {Suit: "heart", Value: 6}, {Suit: "heart", Value: 7}}
	html, _ := template.ParseFiles("main.html")
	html.Execute(writer, test)
}

func web() {
	http.HandleFunc("/game", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

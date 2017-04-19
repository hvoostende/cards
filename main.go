package main

import (
	"fmt"
	"math/rand"
	"time"

	d "github.com/hvoostende/cards/deck"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var deck d.Deck
	deck.BuildDeck()
	// deck.Shuffle()

	for i := 0; i < 3; i++ {
		fmt.Println(deck.Cards[i])
	}

	hand := deck.Deal(3)

	for i := 0; i < 3; i++ {
		fmt.Println(deck.Cards[i])
	}

	fmt.Println(hand)

}

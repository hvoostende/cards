package deck

import (
	"fmt"
	"math/rand"
	"strconv"
)

//Card is the definition of a playing card
type Card struct {
	Suit  string
	Value string
}

//Deck contains a slice of cards
type Deck struct {
	Cards []Card
	Empty bool
}

func (d *Deck) createSuite(suit string) {

	for i := 1; i <= 13; i++ { // create 13 cards
		switch i {
		case 1: // ace
			d.Cards = append(d.Cards, Card{suit, "Ace"})
		case 2, 3, 4, 5, 6, 7, 8, 9, 10: // numbers
			d.Cards = append(d.Cards, Card{suit, strconv.Itoa(i)})
		case 11: // jack
			d.Cards = append(d.Cards, Card{suit, "Jack"})
		case 12: // queen
			d.Cards = append(d.Cards, Card{suit, "Queen"})
		case 13: // king
			d.Cards = append(d.Cards, Card{suit, "King"})
		}
	}
}

//BuildDeck Creates a deck of cards
func (d *Deck) BuildDeck() {

	d.createSuite("Spades")
	d.createSuite("Clubs")
	d.createSuite("Hearts")
	d.createSuite("Diamonds")
	d.Empty = false
}

//Deal removes n cards from the top of the deck and returns the cards in a slice
func (d *Deck) Deal(n int) []Card {
	var hand []Card
	cardsLeft := len(d.Cards)
	if n > cardsLeft {
		fmt.Printf("Dealing last %v cards from deck\n", cardsLeft)

		for i := 0; i < cardsLeft; i++ {
			hand = append(hand, d.Cards[0])
			d.Cards = append(d.Cards[:0], d.Cards[1:]...)
		}
		d.Empty = true
		return hand
	}
	for i := 0; i < n; i++ {
		hand = append(hand, d.Cards[0])
		d.Cards = append(d.Cards[:0], d.Cards[1:]...)
	}
	if len(d.Cards) <= 0 {
		d.Empty = true
	}
	return hand
}

//Shuffle shuffles the deck
func (d *Deck) Shuffle() {
	src := d.Cards
	dest := make([]Card, len(src))
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	d.Cards = dest
}

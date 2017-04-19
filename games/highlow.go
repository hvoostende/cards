package games

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	d "github.com/hvoostende/cards/deck"
)

func value(c d.Card) int {
	switch c.Value {
	case "Ace":
		return 1
	case "2", "3", "4", "5", "6", "7", "8", "9", "10":
		i, err := strconv.Atoi(c.Value)
		if err != nil {
			log.Panic("Something went wrong converting str to int", err)
		}
		return i
	case "Jack":
		return 11
	case "Queen":
		return 12
	case "King":
		return 13
	}
	return 0
}

func HighLow() {
	var deck d.Deck
	deck.BuildDeck()
	fmt.Println("Created a fresh deck..")

	deck.Shuffle()
	fmt.Println("Shuffeling deck..")

	if deck.Empty {
		fmt.Println("unbelievable no more cards left !!")
	}

	scanner := bufio.NewScanner(os.Stdin)
	var name string
	firstCard := true
	var oldCard []d.Card
	fmt.Println("What is your name?")
	for scanner.Scan() {

		if name != "" {
			name = scanner.Text()
		}

		if firstCard {
			oldCard = deck.Deal(1)
			firstCard = false
		}

		fmt.Println(name, "Higher or Lower then:", oldCard[0])
		newCard := deck.Deal(1)
		if scanner.Text() == "higher" || scanner.Text() == "h" || scanner.Text() == "high" {
			if value(oldCard[0]) > value(newCard[0]) {
				fmt.Println(name, "You were right it was:", newCard)
			} else {
				fmt.Println("WROOOOONG!!!")
			}
		}

		if scanner.Text() == "lower" || scanner.Text() == "l" || scanner.Text() == "low" {
			if value(oldCard[0]) < value(newCard[0]) {
				fmt.Println(name, "You were right it was:", newCard)
			} else {
				fmt.Println("WROOOOONG!!!")
			}
		}

		oldCard = newCard

	}

}

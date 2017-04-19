package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hvoostende/cards/games"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Let's Play Higher Lower! :-)")
	games.HighLow()
}

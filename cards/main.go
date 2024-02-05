package main

func main() {
	//card := newCard()
	// fmt.Println(card)

	// cards := deck{newCard(), "Ace of Diamonds"}

	//cards1 := deck{"Test", "Test One"}

	// fmt.Println(cards)

	//cards = append(cards, "Six of Spades")

	// fmt.Println(cards)

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// cards.print()

	// cards1.print()

	cards := newDeck()

	// hand, remain := deal(cards, 10)

	// fmt.Println("Hand")
	// hand.print()

	// fmt.Println("Remainder")
	// remain.print()

	// err := cards.saveToFile("cardfile")
	// if err != nil {
	// 	fmt.Println("An error occurred")
	// }

	// bytes := []byte(cards.toString())
	// cardsNew := string(bytes)
	// fmt.Println(strings.Split(cardsNew, ","))

	// c := newDeckFromFile("cardfile1")
	// c.print()

	cards.shuffle()
	cards.print()

}

// func newCard() string {
// 	return "Five of Diamonds"
// }

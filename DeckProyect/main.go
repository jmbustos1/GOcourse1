package main

import "fmt"

func main() {
	fmt.Println("hi there")
	// var card string = "Ace of Spades"
	// cards := []string{"Ace of Diamonds", newCard()}
	//cards := deck{"Ace of Diamonds", newCard()}

	//cards = append(cards, "Six of Spades")
	cards3 := newDeck()
	cards2 := newDeckFromFile("my_cards")
	cards := newDeck()
	card := newCard()

	cards3.shuffle()

	cards.print()

	hand, remainingCards := deal(cards, 5)
	//for i, car := range cards {
	//	fmt.Println(i, car)
	//}
	greeting := "Hi there!"
	fmt.Println([]byte(greeting))

	hand.print()
	remainingCards.print()

	fmt.Println(card)
	fmt.Println(cards)

	fmt.Println(cards.toString())

	cards.saveToFile("my_cards2")
	fmt.Println(cards2)

	cards3.print()
}

func newCard() string {
	return "Five of Diamonds"
}

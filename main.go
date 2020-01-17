package main

import "fmt"

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Sic of spaced")
	cards.print()
	fmt.Println(cards)

}

func newCard() string {
	return "Five of diamonds"
}

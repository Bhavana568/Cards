//go is not oo lang
//datatypes in go: string integer float array map

//newdeck: craete a list of cards:string
//print: to print
//shuffle: to shuffle cards
//deal: creat a hand of cards
//savetofile
//newdeckfromfile load a list of cards from local machine

package main

//array: fixed length list of things
//slice: an array that cann shrink or grow
func main() {
	// var card string = "Ace of spades" //one method of declaring new var

	//card := newCard() //declaring new var

	//card = "Five of diamonds" //assigning old var new val

	cards := newDeck() //declaring a slice
	cards.shuffle()
	cards.print()

	// hand, remainigCards := deal(cards, 5)

	// hand.print()
	// remainigCards.print()

	// greeting := "Hi there!"
	// fmt.Println(([]byte(greeting)))

}

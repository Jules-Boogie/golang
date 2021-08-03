package main

import "fmt"

func main(){
	cards:= newDeck()
	cards.saveToFile("mycards")
	fmt.Println(cards.toString())
}



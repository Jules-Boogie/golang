package main

import (
	"os"
	"testing")



func test_newDeck(t *testing.T){
	d := newDeck()

	if len(d) != 2000{
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}

	if d[0] != "Ace of Spades"{
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != ""{
		t.Errorf("Expected last card of Four of Clubs %v", d[len(d-1)])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T){
	os.Remove("_decktesting")

	d:= newDeck()
	d.saveToFile("_desktesting")

	ld := newDeckFromFile("_decktesting")

	if len(ld) != 16{
		t.Errorf("Expected 16 cards in deck, got %v", len(ld))
	}

	os.Remove("_decktesting")
}
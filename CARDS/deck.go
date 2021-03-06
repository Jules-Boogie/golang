package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck{
	cards := deck{}
	cardSuits := []string {"Spades", "Diamonds", "hearts","Clubs"}
	cardValues := []string {"Ace", "Two", "Three", "Four"}

	for _, cS := range cardSuits{
		for _, cV := range cardValues{
			cards = append(cards, cV + " of " + cS)
		}
	}
	return cards
}

func (d deck) print(){
	for i,d := range d{
		fmt.Println(i,d)
	}
}

func deal(d deck, handSize int)(deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string{
	return strings.Join([]string(d),",")
}

func (d deck) saveToFile(filename string) error{
return ioutil.WriteFile(filename, []byte(d.toString()),0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile((filename))
	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s:= strings.Split(string(bs),",")
	return deck(s)
}

func (d deck) shuffle(){
	i64 := time.Now().UnixNano()
	source := rand.NewSource(i64)
	r := rand.New(source)
	for i := range d{
		newPosition := r.Intn((len(d)-1))
		d[i], d[newPosition] = d[newPosition],d[i]
	}
}
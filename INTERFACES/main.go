package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printgreeting(eb)
	printgreeting(sb)
}

func printgreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (e englishBot) getGreeting() string {
	return ""
}

func (s spanishBot) getGreeting() string {
	return ""
}

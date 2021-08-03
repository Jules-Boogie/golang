package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	s := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range s {
		go checkStatus(link, c)
	}
	for link := range c{
		go func(l string){
			time.Sleep(5 * time.Second)
			checkStatus(l,c)
		}(link)
	}
}

func checkStatus(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Println(l, "might be down")
		c <- l
		return
	}
	fmt.Println(l, "is up")
	c <- l
}

/*
go scheduler works with one cpu - monitors code running in the differebt go routines

concurrency - we can have multiple threads executing code. a program is concurrent when it has the ability
to run many things kind of at the same time. we can schedule work to be done

parallelism - multiple threads exxecuted at the exact time. requires multiple CPUs

*/

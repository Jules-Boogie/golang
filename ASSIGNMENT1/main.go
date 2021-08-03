package main

import "fmt"

func main() {
	i := []int{0,1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range i {
		fmt.Println(num, isEven(num))
	}
}

func isEven(num int) string {
	if num%2 == 0 {
		return " is even"
	} else {
		return " is odd"
	}
}
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main(){
args := os.Args
readTextFile(args[1])
}

func readTextFile(f string){
	// open text file
	file, err:= os.Open(f)
	if err != nil{
		log.Fatal(err)
		os.Exit(1)
	}
	// read the contents and print to the terminal
	data := make([]byte, 99999)
	file.Read(data)
	fmt.Println(string(data))
}
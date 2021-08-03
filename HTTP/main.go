package main

import (
	"fmt"
	"net/http"
	"os"
)


type logWriter struct {}
func main(){
	resp, err:= http.Get("http://google.com")
	if err != nil{
		fmt.Println("Error", err)
		os.Exit(1)
	}
	lw := logWriter{}
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	io.Copy(os.Stdout, resp.Body)
}

func (l logWriter) Write(bs []byte) (int, error){
	fmt.Println(string(bs))
	return len(bs),nil
}
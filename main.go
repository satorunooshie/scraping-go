package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "http://localhost:8080"

	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	byteArray, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("couldn't read data", err)
		return
	}
	fmt.Println(string(byteArray))
}

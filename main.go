package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/CocoCODEUR/go-scrapping-indeed/scrapping"
)

func main() {

	t := time.Now()

	Jobs := scrapping.Scrapper()
	log.Printf(" job : %v\n", Jobs)
	defer log.Println(time.Since(t))

	data, err := json.MarshalIndent(Jobs, "", " ")
	_ = ioutil.WriteFile("Jobs.json", data, 0644)

	if err != nil {
		log.Printf("Cannot convert data to JSON: %v", err)
	}

}

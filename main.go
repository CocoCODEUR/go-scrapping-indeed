package main

import (
	"log"
	"time"

	"github.com/CocoCODEUR/go-scrapping-indeed/routes"
)

func main() {

	t := time.Now()
	routes.CreateRoutes()
	log.Println(time.Since(t))
}

package main

import (
	"log"

	"github.com/dannypsnl/rocket"
)

func main() {
	rocket.Ignite(":8080").
		Mount("/", rocket.Get("/", func() string {
			log.Println("Receive Request")
			return "serive"
		})).
		Launch()
}

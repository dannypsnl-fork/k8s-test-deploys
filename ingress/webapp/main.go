package main

import (
	"fmt"

	"github.com/dannypsnl/rocket"
)

type Echo struct {
	Message string `route:"message"`
}

func main() {
	rocket.Ignite(":8080").
		Mount("/",
			rocket.Get("/", func() string {
				return "home"
			}),
			rocket.Get("/:message", func(echo *Echo) string {
				fmt.Printf("Heard %s\n", echo.Message)
				return echo.Message
			})).
		Launch()
}

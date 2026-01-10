package main

import (
	"log"

	"github.com/msterhuj/ratioarr/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
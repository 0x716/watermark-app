package main

import (
	"log"

	"github.com/0x716/watermark-app/bootstrap"
)

func main() {
	app := bootstrap.NewApplication()

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

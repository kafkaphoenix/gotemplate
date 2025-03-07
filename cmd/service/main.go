package main

import (
	"log"

	"github.com/kafkaphoenix/gotemplate/cmd/service/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}

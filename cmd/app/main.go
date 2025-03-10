package main

import (
	"log"

	"github.com/kafkaphoenix/gotemplate/cmd/app/bootstrap"

	_ "github.com/kafkaphoenix/gotemplate/docs" // generated docs by swaggo/swag
)

// @title GoTemplate API
// @version 1.0
// @description GoTemplate is a microservice example that follows clean architecture

// @license.name MIT
// @license.url https://github.com/kafkaphoenix/gotemplate/?tab=MIT-1-ov-file#readme

// @contact.name Javier Aguilera
// @contact.email jaguilerapuerta@gmail.com

// @schemes http
// @host localhost:8081
// @BasePath /
func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/akhilbidhuri/shopHere/controller"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	app := controller.App{}
	app.Intitialize()
}

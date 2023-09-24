package main

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/mmvergara/go-simple-api/application"
)

func main(){
	godotenv.Load(".env")
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed to start app:", err)
	}
}
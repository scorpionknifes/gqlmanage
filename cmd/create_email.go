package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	uuid "github.com/nu7hatch/gouuid"

	"github.com/scorpionknifes/gqlmanage/models"
	"github.com/scorpionknifes/gqlmanage/server"
)

// Run this file by using

// go run ./cmd/install.go

func main() {
	// get .env config
	server.GetConfig()

	// connect to db
	collection := server.ConnectDB().Collection("user")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	password, _ := uuid.NewV4()

	user := models.User{
		Name:     "email",
		Username: "email",
		Password: password.String(),
		Location: os.Getenv("DEFAULT_LOCATION"),
		Abbr:     os.Getenv("DEFAULT_ABBR"),
		Email:    os.Getenv("DEFAULT_EMAIL"),
		Openhab:  os.Getenv("DEFAULT_OPENHAB"),
	}

	user.HashPassword(user.Password)
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Panic()
	}
	fmt.Println("CREATED USER email WITH PASSWORD:")
	fmt.Println(password)
}

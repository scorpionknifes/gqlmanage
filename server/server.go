package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/scorpionknifes/gqlopenhab/dataloader"
	"github.com/scorpionknifes/gqlopenhab/graphql"
	"github.com/scorpionknifes/gqlopenhab/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const defaultPort = "8080"

func main() {
	// Connect to MongoDB

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost/"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("tsatech")

	c := graphql.Config{Resolvers: &graphql.Resolver{
		UserRepo:   mongodb.UserRepo{DB: db.Collection("user")},
		DeviceRepo: mongodb.DeviceRepo{DB: db.Collection("device")},
		RoomRepo:   mongodb.RoomRepo{DB: db.Collection("room")},
	}}

	d := &dataloader.DBLoader{
		RoomRepo:   mongodb.RoomRepo{DB: db.Collection("room")},
		DeviceRepo: mongodb.DeviceRepo{DB: db.Collection("device")},
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", dataloader.DataMiddleware(d, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

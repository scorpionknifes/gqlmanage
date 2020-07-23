package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	"github.com/scorpionknifes/gqlopenhab/dataloader"
	"github.com/scorpionknifes/gqlopenhab/graphql"
	customMiddleware "github.com/scorpionknifes/gqlopenhab/middleware"
	"github.com/scorpionknifes/gqlopenhab/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const defaultPort = "8080"

func main() {
	// Connect to MongoDB

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(os.Getenv("MONGODB_DATABASE"))
	var (
		userRepo   = mongodb.UserRepo{DB: db.Collection("user")}
		deviceRepo = mongodb.DeviceRepo{DB: db.Collection("device")}
		roomRepo   = mongodb.RoomRepo{DB: db.Collection("room")}
	)

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8000"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(customMiddleware.AuthMiddleware(userRepo))

	c := graphql.Config{Resolvers: &graphql.Resolver{
		UserRepo:   userRepo,
		DeviceRepo: deviceRepo,
		RoomRepo:   roomRepo,
	}}

	d := &dataloader.DBLoader{
		DeviceRepo: deviceRepo,
		RoomRepo:   roomRepo,
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

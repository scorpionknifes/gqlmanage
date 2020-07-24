package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	"github.com/scorpionknifes/gqlmanage/dataloader"
	"github.com/scorpionknifes/gqlmanage/graphql"
	customMiddleware "github.com/scorpionknifes/gqlmanage/middleware"
	"github.com/scorpionknifes/gqlmanage/mongodb"
)

const defaultPort = "8080"

func main() {
	db := connectDB()

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

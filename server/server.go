package server

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
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

// Init start server
func Init() {
	// get .env config
	getConfig()

	// connect to db
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

	d := &dataloader.DBLoader{
		DeviceRepo: deviceRepo,
		RoomRepo:   roomRepo,
	}
	router.Use(dataloader.DataMiddleware(d))

	c := graphql.Config{Resolvers: &graphql.Resolver{
		UserRepo:   userRepo,
		DeviceRepo: deviceRepo,
		RoomRepo:   roomRepo,
	}}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.New(graphql.NewExecutableSchema(c))
	srv.AddTransport(transport.POST{})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

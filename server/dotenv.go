package server

import "github.com/joho/godotenv"

func getConfig() {
	godotenv.Load()
}

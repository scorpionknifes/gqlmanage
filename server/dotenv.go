package server

import "github.com/joho/godotenv"

// GetConfig by using godotenv
func GetConfig() {
	godotenv.Load()
}

//go:generate go run github.com/99designs/gqlgen

package graphql

import (
	"github.com/go-redis/redis/v8"
	"github.com/scorpionknifes/gqlmanage/mongodb"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver struct
type Resolver struct {
	DeviceRepo mongodb.DeviceRepo
	RoomRepo   mongodb.RoomRepo
	UserRepo   mongodb.UserRepo
	EmailRepo  mongodb.EmailRepo
	Redis      redis.Client
}

package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type RoomRepo struct {
	DB *mongo.Collection
}

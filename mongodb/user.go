package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type UserRepo struct {
	DB *mongo.Collection
}

package mongodb

import (
	"context"
	"time"

	"github.com/scorpionknifes/gqlopenhab/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepo Collection
type UserRepo struct {
	DB *mongo.Collection
}

// GetUsers get all users
func (d *UserRepo) GetUsers() ([]*models.User, error) {
	var users []*models.User

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := d.DB.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

// GetUsersByRoomID get all users in a Room by roomID
func (d *UserRepo) GetUsersByRoomID(roomID string) ([]*models.User, error) {
	var users []*models.User

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := d.DB.Find(ctx, bson.M{"room": roomID})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUser create user
func (d *UserRepo) CreateUser(roomID primitive.ObjectID, user *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := d.DB.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

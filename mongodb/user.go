package mongodb

import (
	"context"
	"time"

	"github.com/scorpionknifes/gqlmanage/models"
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

// GetUser get user by id
func (d *UserRepo) GetUser(id string) (*models.User, error) {
	var user *models.User

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result := d.DB.FindOne(ctx, bson.M{"_id": ID})
	err = result.Decode(&user)
	return user, err
}

// GetUserByUsername get user by username
func (d *UserRepo) GetUserByUsername(username string) (*models.User, error) {
	var user *models.User

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result := d.DB.FindOne(ctx, bson.M{"username": username})
	err := result.Decode(&user)
	return user, err
}

// CreateUser create user
func (d *UserRepo) CreateUser(user *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	user.HashPassword(user.Password)
	result, err := d.DB.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errNoID
	}
	user.ID = oid.Hex()
	return user, nil
}

// UpdateUser update a User by id
func (d *UserRepo) UpdateUser(id string, user *models.User) (*models.User, error) {
	if user.Password != "" {
		user.HashPassword(user.Password)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	_, err = d.DB.UpdateOne(ctx, bson.M{"_id": ID}, bson.M{"$set": user})
	if err != nil {
		return nil, err
	}
	user.ID = id
	return user, nil
}

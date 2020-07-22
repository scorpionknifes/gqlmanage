package mongodb

import (
	"context"
	"errors"
	"time"

	"github.com/scorpionknifes/gqlopenhab/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// RoomRepo Collection
type RoomRepo struct {
	DB *mongo.Collection
}

// GetRooms get all rooms
func (d *RoomRepo) GetRooms() ([]*models.Room, error) {
	var rooms []*models.Room

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := d.DB.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &rooms); err != nil {
		return nil, err
	}
	return rooms, nil
}

// GetRoom get one room
func (d *RoomRepo) GetRoom(id string) (*models.Room, error) {
	var room *models.Room

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result := d.DB.FindOne(ctx, bson.M{"_id": ID})
	err = result.Decode(room)
	return room, err
}

// CreateRoom create room
func (d *RoomRepo) CreateRoom(room *models.Room) (*models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, err := d.DB.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("Bad ID")
	}
	room.ID = oid.Hex()
	return room, nil
}

// UpdateRoom update a room by id
func (d *RoomRepo) UpdateRoom(id string, room *models.Room) (*models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := d.DB.UpdateOne(ctx, bson.M{"_id": id}, room)
	if err != nil {
		return nil, err
	}
	return room, nil
}

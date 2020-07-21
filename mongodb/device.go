package mongodb

import (
	"context"
	"time"

	"github.com/scorpionknifes/gqlopenhab/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DeviceRepo struct {
	DB *mongo.Collection
}

func (d *DeviceRepo) GetDevicesByRoomID(roomID string) ([]*models.Device, error) {
	var devices []*models.Device

	id, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		return devices, err
	}
	q := bson.M{"room": id}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := d.DB.Find(ctx, q, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var device *models.Device
		err = cur.Decode(&device)
		if err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(ctx)
	return devices, nil
}

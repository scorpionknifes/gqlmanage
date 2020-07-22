package mongodb

import (
	"context"
	"time"

	"github.com/scorpionknifes/gqlopenhab/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// DeviceRepo Collection
type DeviceRepo struct {
	DB *mongo.Collection
}

// GetDevices get all devices
func (d *DeviceRepo) GetDevices() ([]*models.Device, error) {
	var devices []*models.Device

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := d.DB.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &devices); err != nil {
		return nil, err
	}
	return devices, nil
}

// GetDevice get one device
func (d *DeviceRepo) GetDevice(id string) (*models.Device, error) {
	var device *models.Device

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result := d.DB.FindOne(ctx, bson.M{"_id": ID})
	err = result.Decode(device)
	return device, err
}

// GetDevicesByDeviceID get all devices in a Device by deviceID
func (d *DeviceRepo) GetDevicesByDeviceID(deviceID string) ([]*models.Device, error) {
	var devices []*models.Device

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := d.DB.Find(ctx, bson.M{"room_id": deviceID})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &devices); err != nil {
		return nil, err
	}
	return devices, nil
}

// CreateDevice create device
func (d *DeviceRepo) CreateDevice(device *models.Device) (*models.Device, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := d.DB.InsertOne(ctx, device)
	if err != nil {
		return nil, err
	}
	return device, nil
}

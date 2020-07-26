package mongodb

import (
	"context"
	"time"

	"github.com/scorpionknifes/gqlmanage/models"
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
	err = result.Decode(&device)
	return device, err
}

// GetDevicesByRoomID get all devices in a Device by deviceID
func (d *DeviceRepo) GetDevicesByRoomID(deviceID string) ([]*models.Device, error) {
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
	result, err := d.DB.InsertOne(ctx, device)
	if err != nil {
		return nil, err
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errNoID
	}
	device.ID = oid.Hex()
	return device, nil
}

// UpdateDevice update a device by id
func (d *DeviceRepo) UpdateDevice(id string, device *models.Device) (*models.Device, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	_, err = d.DB.UpdateOne(ctx, bson.M{"_id": ID}, bson.M{"$set": device})
	if err != nil {
		return nil, err
	}
	device.ID = id
	return device, nil
}

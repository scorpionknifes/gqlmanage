package graphql

import (
	"context"

	"github.com/scorpionknifes/gqlopenhab/models"
)

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return r.UserRepo.GetUsers()
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return r.UserRepo.GetUser(id)
}

func (r *queryResolver) Rooms(ctx context.Context) ([]*models.Room, error) {
	return r.RoomRepo.GetRooms()
}

func (r *queryResolver) Room(ctx context.Context, id string) (*models.Room, error) {
	return r.RoomRepo.GetRoom(id)
}

func (r *queryResolver) Devices(ctx context.Context) ([]*models.Device, error) {
	return r.DeviceRepo.GetDevices()
}

func (r *queryResolver) Device(ctx context.Context, id string) (*models.Device, error) {
	return r.DeviceRepo.GetDevice(id)
}

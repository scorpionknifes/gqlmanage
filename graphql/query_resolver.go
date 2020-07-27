package graphql

import (
	"context"

	"github.com/scorpionknifes/gqlmanage/middleware"
	"github.com/scorpionknifes/gqlmanage/models"
)

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}
	return r.UserRepo.GetUsers()
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}
	return r.UserRepo.GetUser(id)
}

func (r *queryResolver) Rooms(ctx context.Context, filter *models.RoomFilter, limit *int, offset *int) ([]*models.Room, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}
	return r.RoomRepo.GetRooms(filter, limit, offset)
}

func (r *queryResolver) Room(ctx context.Context, id string) (*models.Room, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}
	return r.RoomRepo.GetRoom(id)
}

func (r *queryResolver) Devices(ctx context.Context) ([]*models.Device, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}
	return r.DeviceRepo.GetDevices()
}

func (r *queryResolver) Device(ctx context.Context, id string) (*models.Device, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}
	return r.DeviceRepo.GetDevice(id)
}

func (r *queryResolver) Email(ctx context.Context, id string) (*models.Email, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}
	return r.EmailRepo.GetEmail(id)
}

func (r *queryResolver) Emails(ctx context.Context) ([]*models.Email, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}
	return r.EmailRepo.GetEmails()
}

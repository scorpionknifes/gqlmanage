//go:generate go run github.com/99designs/gqlgen

package graphql

import (
	"context"
	"fmt"

	"github.com/scorpionknifes/gqlopenhab/models"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver struct
type Resolver struct {
}

func (r *mutationResolver) Login(ctx context.Context, email *string, password *string) (*Token, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateRoom(ctx context.Context, input *CreateRoomInput) (*models.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddDevice(ctx context.Context, input *AddDeviceInput) (*models.Device, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Rooms(ctx context.Context) ([]*models.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Room(ctx context.Context, id string) (*models.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Devices(ctx context.Context) ([]*models.Device, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Device(ctx context.Context, id string) (*models.Device, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *deviceResolver) Room(ctx context.Context, obj *models.Device) (*models.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *roomResolver) Devices(ctx context.Context, obj *models.Room) ([]*models.Device, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Device returns DeviceResolver implementation.
func (r *Resolver) Device() DeviceResolver { return &deviceResolver{r} }

// Room returns RoomResolver implementaion.
func (r *Resolver) Room() RoomResolver { return &roomResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

type deviceResolver struct{ *Resolver }
type roomResolver struct{ *Resolver }

//go:generate go run github.com/99designs/gqlgen

package graphql

import (
	"context"
	"fmt"
	"time"

	"github.com/scorpionknifes/gqlopenhab/models"
	"github.com/scorpionknifes/gqlopenhab/mongodb"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver struct
type Resolver struct {
	DeviceRepo mongodb.DeviceRepo
	RoomRepo   mongodb.RoomRepo
	UserRepo   mongodb.UserRepo
}

func (r *mutationResolver) Login(ctx context.Context, input LoginInput) (*Token, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateRoom(ctx context.Context, input RoomInput) (*models.Room, error) {
	room := &models.Room{
		RoomNumber:  input.RoomNumber,
		Memo:        input.Memo,
		Username:    input.Username,
		Password:    input.Password,
		CreatedDate: time.Now(),
	}
	return r.RoomRepo.CreateRoom(room)
}

func (r *mutationResolver) UpdateRoom(ctx context.Context, id string, input RoomInput) (*models.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateDevice(ctx context.Context, input DeviceInput) (*models.Device, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDevice(ctx context.Context, id string, input DeviceInput) (*models.Device, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input UserInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input UserInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

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

func (r *deviceResolver) Room(ctx context.Context, obj *models.Device) (*models.Room, error) {
	return r.RoomRepo.GetRoom(obj.RoomID)
}

func (r *roomResolver) Devices(ctx context.Context, obj *models.Room) ([]*models.Device, error) {
	return r.DeviceRepo.GetDevicesByDeviceID(obj.ID)
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

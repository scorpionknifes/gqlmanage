package graphql

import (
	"context"
	"fmt"
	"time"

	"github.com/scorpionknifes/gqlopenhab/models"
)

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

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
	room := &models.Room{
		RoomNumber:  input.RoomNumber,
		Memo:        input.Memo,
		Username:    input.Username,
		Password:    input.Password,
		CreatedDate: time.Now(),
	}
	return r.RoomRepo.UpdateRoom(id, room)
}

func (r *mutationResolver) CreateDevice(ctx context.Context, input DeviceInput) (*models.Device, error) {
	device := &models.Device{
		RoomID:       input.RoomID,
		Name:         input.Name,
		Model:        input.Model,
		MacAddress:   input.MacAddress,
		Memo:         input.Memo,
		SerialNumber: input.SerialNumber,
		Status:       input.Status,
		Type:         input.Type,
	}
	return r.DeviceRepo.CreateDevice(device)
}

func (r *mutationResolver) UpdateDevice(ctx context.Context, id string, input DeviceInput) (*models.Device, error) {
	device := &models.Device{
		RoomID:       input.RoomID,
		Name:         input.Name,
		Model:        input.Model,
		MacAddress:   input.MacAddress,
		Memo:         input.Memo,
		SerialNumber: input.SerialNumber,
		Status:       input.Status,
		Type:         input.Type,
	}
	return r.DeviceRepo.UpdateDevice(id, device)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input UserInput) (*models.User, error) {
	user := &models.User{
		Name:     input.Name,
		Username: input.Username,
		Password: input.Password,
		Location: input.Location,
		Abbr:     input.Abbr,
		Email:    input.Email,
		Openhab:  input.Openhab,
	}
	return r.UserRepo.CreateUser(user)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input UserInput) (*models.User, error) {
	user := &models.User{
		Name:     input.Name,
		Username: input.Username,
		Password: input.Password,
		Location: input.Location,
		Abbr:     input.Abbr,
		Email:    input.Email,
		Openhab:  input.Openhab,
	}
	return r.UserRepo.UpdateUser(id, user)
}

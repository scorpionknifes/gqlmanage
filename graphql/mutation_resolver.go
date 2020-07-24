package graphql

import (
	"context"
	"time"

	"github.com/scorpionknifes/gqlmanage/models"
)

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	user, err := r.UserRepo.GetUserByUsername(input.Username)

	if err != nil {
		return nil, errBadCredentials
	}

	err = user.ComparePassword(input.Password)
	if err != nil {
		return nil, errBadCredentials
	}

	token, err := user.GenToken()
	if err != nil {
		return nil, errUnknown
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (r *mutationResolver) CreateRoom(ctx context.Context, input models.RoomInput) (*models.Room, error) {
	room := &models.Room{
		RoomNumber:  input.RoomNumber,
		Memo:        input.Memo,
		Username:    input.Username,
		Password:    input.Password,
		CreatedDate: time.Now(),
	}
	return r.RoomRepo.CreateRoom(room)
}

func (r *mutationResolver) UpdateRoom(ctx context.Context, id string, input models.RoomUpdate) (*models.Room, error) {
	room := &models.Room{}
	room.Update(input)
	return r.RoomRepo.UpdateRoom(id, room)
}

func (r *mutationResolver) CreateDevice(ctx context.Context, input models.DeviceInput) (*models.Device, error) {
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

func (r *mutationResolver) UpdateDevice(ctx context.Context, id string, input models.DeviceUpdate) (*models.Device, error) {
	device := &models.Device{}
	device.Update(input)
	return r.DeviceRepo.UpdateDevice(id, device)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
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

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input models.UserUpdate) (*models.User, error) {
	user := &models.User{}
	user.Update(input)
	return r.UserRepo.UpdateUser(id, user)
}

package graphql

import (
	"context"
	"encoding/json"
	"time"

	"github.com/scorpionknifes/gqlmanage/middleware"
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
	_, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}
	room := &models.Room{
		RoomNumber:  input.RoomNumber,
		Memo:        input.Memo,
		Username:    input.Username,
		Password:    input.Password,
		CreatedDate: time.Now(),
	}
	room, err = r.RoomRepo.CreateRoom(room)
	if err != nil {
		return nil, err
	}
	for _, device := range input.Devices {
		d := &models.Device{
			RoomID:       room.ID,
			Name:         device.Name,
			Model:        device.Model,
			MacAddress:   device.MacAddress,
			Memo:         device.Memo,
			SerialNumber: device.SerialNumber,
			Status:       device.Status,
			Type:         device.Type,
			CreatedDate:  time.Now(),
			LastModified: time.Now(),
		}
		r.DeviceRepo.CreateDevice(d)
	}
	return room, nil
}

func (r *mutationResolver) UpdateRoom(ctx context.Context, id string, input models.RoomUpdate) (*models.Room, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}
	room, err := r.RoomRepo.GetRoom(id)
	if err != nil {
		return nil, err
	}
	room.ID = ""
	room.Update(input)
	return r.RoomRepo.UpdateRoom(id, room)
}

func (r *mutationResolver) CreateDevice(ctx context.Context, input models.DeviceInput) (*models.Device, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}
	device := &models.Device{
		RoomID:       input.RoomID,
		Name:         input.Name,
		Model:        input.Model,
		MacAddress:   input.MacAddress,
		Memo:         input.Memo,
		SerialNumber: input.SerialNumber,
		Status:       input.Status,
		Type:         input.Type,
		CreatedDate:  time.Now(),
		LastModified: time.Now(),
	}
	return r.DeviceRepo.CreateDevice(device)
}

func (r *mutationResolver) UpdateDevice(ctx context.Context, id string, input models.DeviceUpdate) (*models.Device, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}
	device, err := r.DeviceRepo.GetDevice(id)
	if err != nil {
		return nil, err
	}
	device.ID = ""
	device.Update(input)
	return r.DeviceRepo.UpdateDevice(id, device)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}
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
	_, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}
	user, err := r.UserRepo.GetUser(id)
	if err != nil {
		return nil, err
	}
	user.Password = ""
	user.ID = ""
	user.Update(input)
	return r.UserRepo.UpdateUser(id, user)
}

func (r *mutationResolver) CreateEmail(ctx context.Context, input models.EmailInput) (*models.Email, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}
	email := &models.Email{
		From:        input.From,
		To:          input.To,
		Data:        input.Data,
		CreatedDate: time.Now(),
	}

	email, err = r.EmailRepo.CreateEmail(email)
	if err != nil {
		return nil, err
	}
	emailJSON, err := json.Marshal(email)
	if err != nil {
		return nil, err
	}
	err = r.Redis.Publish(ctx, "email", emailJSON).Err()
	if err != nil {
		return nil, err
	}
	return email, nil
}

package graphql

import (
	"context"

	"github.com/scorpionknifes/gqlmanage/dataloader"
	"github.com/scorpionknifes/gqlmanage/models"
)

// Device returns DeviceResolver implementation.
func (r *Resolver) Device() DeviceResolver { return &deviceResolver{r} }

type deviceResolver struct{ *Resolver }

func (r *deviceResolver) Room(ctx context.Context, obj *models.Device) (*models.Room, error) {
	return dataloader.GetRoomLoader(ctx).Load(obj.RoomID)
}

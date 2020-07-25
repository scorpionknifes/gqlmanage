package dataloader

import (
	"context"
	"net/http"
	"time"

	"github.com/scorpionknifes/gqlmanage/models"
	"github.com/scorpionknifes/gqlmanage/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type key string

const (
	deviceloaderKey key = "deviceloader"
	roomloaderKey   key = "roomloader"
)

// DBLoader for middleware
type DBLoader struct {
	DeviceRepo mongodb.DeviceRepo
	RoomRepo   mongodb.RoomRepo
}

// DataMiddleware to dataloader middleware
func DataMiddleware(db *DBLoader) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			deviceloader := DeviceLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(ids []string) ([]*models.Device, []error) {
					var devices []*models.Device
					oids := make([]primitive.ObjectID, len(ids))
					for i, id := range ids {
						oids[i], _ = primitive.ObjectIDFromHex(id)
					}
					ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
					defer cancel()
					cursor, err := db.DeviceRepo.DB.Find(ctx, bson.M{"_id": bson.M{"$in": oids}})
					if err != nil {
						return nil, []error{err}
					}
					if err := cursor.All(ctx, &devices); err != nil {
						return nil, []error{err}
					}
					d := make(map[string]*models.Device, len(ids))

					for _, device := range devices {
						d[device.ID] = device
					}

					result := make([]*models.Device, len(ids))

					for i, id := range ids {
						result[i] = d[id]
					}

					return result, nil
				},
			}
			roomloader := RoomLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(ids []string) ([]*models.Room, []error) {
					var rooms []*models.Room
					oids := make([]primitive.ObjectID, len(ids))
					for i, id := range ids {
						oids[i], _ = primitive.ObjectIDFromHex(id)
					}
					ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
					defer cancel()
					cursor, err := db.RoomRepo.DB.Find(ctx, bson.M{"_id": bson.M{"$in": oids}})
					if err != nil {
						return nil, []error{err}
					}
					if err := cursor.All(ctx, &rooms); err != nil {
						return nil, []error{err}
					}
					r := make(map[string]*models.Room, len(ids))

					for _, room := range rooms {
						r[room.ID] = room
					}

					result := make([]*models.Room, len(ids))

					for i, id := range ids {
						result[i] = r[id]
					}

					return result, nil
				},
			}
			ctx := context.WithValue(r.Context(), deviceloaderKey, &deviceloader)
			ctx = context.WithValue(ctx, roomloaderKey, &roomloader)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetDeviceLoader get device dataloader
func GetDeviceLoader(ctx context.Context) *DeviceLoader {
	return ctx.Value(deviceloaderKey).(*DeviceLoader)
}

// GetRoomLoader get device dataloader
func GetRoomLoader(ctx context.Context) *RoomLoader {
	return ctx.Value(roomloaderKey).(*RoomLoader)
}

package dataloader

import (
	"context"
	"net/http"
	"time"

	"github.com/scorpionknifes/gqlopenhab/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type key string

const (
	deviceloaderKey key = "deviceloader"
	roomloaderKey   key = "roomloader"
)

// DeviceMiddleware dataloader Middleware
func DeviceMiddleware(db *mongo.Collection, next http.Handler) http.Handler {
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
				cursor, err := db.Find(ctx, bson.M{"_id": bson.M{"$in": oids}})
				if err != nil {
					return nil, []error{err}
				}
				if err := cursor.All(ctx, &devices); err != nil {
					return nil, []error{err}
				}
				return devices, nil
			},
		}

		ctx := context.WithValue(r.Context(), deviceloaderKey, &deviceloader)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RoomMiddleware dataloader Middleware
func RoomMiddleware(db *mongo.Collection, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
				cursor, err := db.Find(ctx, bson.M{"_id": bson.M{"$in": oids}})
				if err != nil {
					return nil, []error{err}
				}
				if err := cursor.All(ctx, &rooms); err != nil {
					return nil, []error{err}
				}
				return rooms, nil
			},
		}
		ctx := context.WithValue(r.Context(), roomloaderKey, &roomloader)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

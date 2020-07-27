package mongodb

import (
	"context"
	"time"

	"github.com/scorpionknifes/gqlmanage/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// EmailRepo Collection
type EmailRepo struct {
	DB *mongo.Collection
}

// GetEmails get all emails
func (d *EmailRepo) GetEmails() ([]*models.Email, error) {
	var emails []*models.Email

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := d.DB.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &emails); err != nil {
		return nil, err
	}
	return emails, nil
}

// GetEmail get one email
func (d *EmailRepo) GetEmail(id string) (*models.Email, error) {
	var email *models.Email

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result := d.DB.FindOne(ctx, bson.M{"_id": ID})
	err = result.Decode(&email)
	return email, err
}

// CreateEmail create email
func (d *EmailRepo) CreateEmail(email *models.Email) (*models.Email, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, err := d.DB.InsertOne(ctx, email)
	if err != nil {
		return nil, err
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errNoID
	}
	email.ID = oid.Hex()
	return email, nil
}

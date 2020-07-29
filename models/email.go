package models

import "time"

// Email struct for email
type Email struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	From        string    `json:"from" bson:"from,omitempty"`
	To          string    `json:"to" bson:"to,omitempty"`
	Data        string    `json:"data" bson:"data,omitempty"`
	CreatedDate time.Time `json:"createdDate" bson:"created_date,omitempty"`
}

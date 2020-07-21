package models

import "time"

// Device struct for one device
type Device struct {
	ID           string    `json:"id" bson:"_id,omitempty"`
	RoomID       string    `bson:"room_id"`
	Name         string    `json:"name" bson:"name"`
	Model        string    `json:"model" bson:"model"`
	MacAddress   string    `json:"macAddress" bson:"mac_address"`
	Memo         string    `json:"memo" bson:"memo"`
	SerialNumber string    `json:"serialNumber" bson:"serial_number"`
	Status       int       `json:"status" bson:"status"`
	Type         int       `json:"type" bson:"type"`
	CreatedDate  time.Time `json:"createdDate" bson:"type"`
	LastModified time.Time `json:"lastModified" bson:"created_date"`
}

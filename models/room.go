package models

import "time"

// Room struct for one room
type Room struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	RoomNumber  string    `json:"roomNumber" bson:"room_number"`
	Memo        string    `json:"memo" bson:"memo"`
	Username    string    `json:"username" bson:"username"`
	Password    string    `json:"password" bson:"password"`
	CreatedDate time.Time `json:"createdDate" bson:"created_date"`
}

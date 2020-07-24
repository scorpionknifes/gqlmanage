package models

import "time"

// Room struct for one room
type Room struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	RoomNumber  string    `json:"roomNumber" bson:"room_number,omitempty"`
	Memo        string    `json:"memo" bson:"memo,omitempty"`
	Username    string    `json:"username" bson:"username,omitempty"`
	Password    string    `json:"password" bson:"password,omitempty"`
	CreatedDate time.Time `json:"createdDate" bson:"created_date,omitempty"`
}

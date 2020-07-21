package models

import "time"

type Room struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	RoomNumber  string    `json:"roomNumber" bson:"room_number"`
	Memo        string    `json:"memo" bson:"memo"`
	Devices     []*Device `json:"devices"`
	Username    string    `json:"username" bson:"username"`
	Password    string    `json:"password" bson:"password"`
	CreatedDate time.Time `json:"createdDate" bson:"created_date"`
}

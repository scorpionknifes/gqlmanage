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

// Update convert RoomUpdate to struct
func (d *Room) Update(input RoomUpdate) {
	if input.RoomNumber != nil {
		d.RoomNumber = *input.RoomNumber
	}
	if input.Memo != nil {
		d.Memo = *input.Memo
	}
	if input.Username != nil {
		d.Username = *input.Username
	}
	if input.Password != nil {
		d.Password = *input.Password
	}
}

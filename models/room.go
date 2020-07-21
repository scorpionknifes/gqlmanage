package models

import "time"

type Room struct {
	ID          string    `json:"id"`
	RoomNumber  string    `json:"roomNumber"`
	Memo        string    `json:"memo"`
	Devices     []*Device `json:"devices"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	CreatedDate time.Time `json:"createdDate"`
}

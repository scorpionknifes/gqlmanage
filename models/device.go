package models

import "time"

type Device struct {
	ID           string    `json:"id"`
	Room         *Room     `json:"room"`
	Name         string    `json:"name"`
	Model        string    `json:"model"`
	MacAddress   string    `json:"macAddress"`
	Memo         string    `json:"memo"`
	SerialNumber string    `json:"serialNumber"`
	Status       int       `json:"status"`
	CreatedDate  time.Time `json:"createdDate"`
	LastModified time.Time `json:"lastModified"`
}

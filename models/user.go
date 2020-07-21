package models

type User struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Abbr     string `json:"abbr"`
	Email    string `json:"email"`
	Openhab  string `json:"openhab"`
}

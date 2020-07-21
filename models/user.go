package models

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Location string `json:"location"`
	Abbr     string `json:"abbr"`
	Email    string `json:"email"`
	Openhab  string `json:"openhab"`
}

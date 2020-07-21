package models

// User struct for one user
type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Name     string `json:"name" bson:"name"`
	Username string `json:"username" bson:"username"`
	Location string `json:"location" bson:"location"`
	Abbr     string `json:"abbr" bson:"abbr"`
	Email    string `json:"email" bson:"email"`
	Openhab  string `json:"openhab" bson:"openhab"`
}

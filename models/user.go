package models

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// User struct for one user
type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Name     string `json:"name" bson:"name,omitempty"`
	Username string `json:"username" bson:"username,omitempty"`
	Password string `json:"password" bson:"password,omitempty"`
	Location string `json:"location" bson:"location,omitempty"`
	Abbr     string `json:"abbr" bson:"abbr,omitempty"`
	Email    string `json:"email" bson:"email,omitempty"`
	Openhab  string `json:"openhab" bson:"openhab,omitempty"`
}

// HashPassword using bcrypt
func (u *User) HashPassword(password string) error {
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(passwordHash)

	return nil
}

// ComparePassword compares password using bcrypt
func (u *User) ComparePassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

// GenToken generate JWT token
func (u *User) GenToken() (*AuthToken, error) {
	expiredAt := time.Now().Add(time.Hour * 24 * 7) // a week

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expiredAt.Unix(),
		Id:        u.ID,
		IssuedAt:  time.Now().Unix(),
		Issuer:    os.Getenv("JWT_ISSUER"),
	})

	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &AuthToken{
		AccessToken: accessToken,
		ExpiredAt:   expiredAt,
	}, nil
}

// Update convert UserUpdate to struct
func (u *User) Update(input UserUpdate) {
	if input.Name != nil {
		u.Name = *input.Name
	}
	if input.Username != nil {
		u.Username = *input.Username
	}
	if input.Password != nil {
		u.Password = *input.Password
	}
	if input.Location != nil {
		u.Location = *input.Location
	}
	if input.Abbr != nil {
		u.Abbr = *input.Abbr
	}
	if input.Email != nil {
		u.Email = *input.Email
	}
	if input.Openhab != nil {
		u.Openhab = *input.Openhab
	}
}

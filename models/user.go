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
	Name     string `json:"name" bson:"name"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Location string `json:"location" bson:"location"`
	Abbr     string `json:"abbr" bson:"abbr"`
	Email    string `json:"email" bson:"email"`
	Openhab  string `json:"openhab" bson:"openhab"`
}

func (u *User) HashPassword(password string) error {
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(passwordHash)

	return nil
}

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

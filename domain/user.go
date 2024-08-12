package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	UserID int    `json:"userID"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

type UserRepository interface {
	CreateUser(user *User) error
	AuthenticateUser(username, password string) (string, error)
	GetUserByUsername(username string) (*User, error)
	ValidateToken(tokenString string) (*Claims, error)
}

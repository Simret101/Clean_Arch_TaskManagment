package repository

import (
	"errors"
	"fmt"
	"sync"
	"task/config"
	"task/domain"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type InMemoryUserRepository struct {
	users      []domain.User
	lastUserID int
	userMu     sync.Mutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users:      []domain.User{},
		lastUserID: 0,
	}
}

func (r *InMemoryUserRepository) CreateUser(user *domain.User) error {
	r.userMu.Lock()
	defer r.userMu.Unlock()

	for _, u := range r.users {
		if u.Username == user.Username {
			return errors.New("username already exists")
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	r.lastUserID++
	user.ID = r.lastUserID

	r.users = append(r.users, *user)
	fmt.Printf("User created: %v\n", user)
	return nil
}

func (r *InMemoryUserRepository) AuthenticateUser(username, password string) (string, error) {
	r.userMu.Lock()
	defer r.userMu.Unlock()

	for _, user := range r.users {
		if user.Username == username {
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err == nil {
				return r.generateToken(user)
			}
			break
		}
	}
	return "", errors.New("invalid username or password")
}

func (r *InMemoryUserRepository) GetUserByUsername(username string) (*domain.User, error) {
	r.userMu.Lock()
	defer r.userMu.Unlock()

	for _, user := range r.users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *InMemoryUserRepository) generateToken(user domain.User) (string, error) {
	claims := domain.Claims{
		UserID: user.ID,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.TokenExpiration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (r *InMemoryUserRepository) ValidateToken(tokenString string) (*domain.Claims, error) {
	claims := &domain.Claims{}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if !token.Valid {
		return nil, errors.New("token is not valid")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}

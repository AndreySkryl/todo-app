package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/AndreySkryl/todo-app"
	"github.com/AndreySkryl/todo-app/pkg/repository"
)

const salt = "ub9uh91uncu912ybv332csfasfqa"

type AuthService struct {
	repository repository.Authorization
}

func NewAuthService(repository repository.Authorization) *AuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repository.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

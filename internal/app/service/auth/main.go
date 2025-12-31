package auth

import (
	"context"

	"github.com/sugandasu/go-boilerplate/config"
	"github.com/sugandasu/go-boilerplate/internal/app/model"
	"github.com/sugandasu/go-boilerplate/internal/app/repository"
)

type AuthService interface {
	Login(ctx context.Context, data model.LoginRequest) (*model.LoginResponse, error)
}

type service struct {
	config         *config.Config
	userRepository repository.UserRepository
}

func NewAuthService(config *config.Config, userRepository repository.UserRepository) AuthService {
	if config == nil {
		panic("config is nil")
	}
	if userRepository == nil {
		panic("userRepository is nil")
	}

	return &service{
		config:         config,
		userRepository: userRepository,
	}
}

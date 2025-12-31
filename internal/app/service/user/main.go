package user

import (
	"context"

	"github.com/sugandasu/go-boilerplate/internal/app/model"
	"github.com/sugandasu/go-boilerplate/internal/app/repository"
	"github.com/sugandasu/ruru/tolo"
)

type UserService interface {
	CreateUser(ctx context.Context, data model.UserCreateRequest) (*model.User, error)
	ListUser(ctx context.Context, filters model.UserFilter) (tolo.PaginationResponse, error)
	UpdateUser(ctx context.Context, data model.UserUpdateRequest) error
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	DeleteUser(ctx context.Context, id string) error
}

type service struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	if userRepo == nil {
		panic("userRepo is nil")
	}

	return &service{
		userRepo: userRepo,
	}
}

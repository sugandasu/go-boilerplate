package user

import (
	"context"
	"net/http"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/sugandasu/go-boilerplate/internal/app/model"
	"github.com/sugandasu/ruru/tolo"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) CreateUser(ctx context.Context, data model.UserCreateRequest) (*model.User, error) {
	if data.RoleID != nil {
		_, err := s.userRepo.GetRoleByID(ctx, *data.RoleID)
		if err != nil {
			return nil, tolo.NewError(http.StatusBadRequest, "fail to get role", map[string]string{
				"role_id": "role is not found",
			})
		}
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 0)

	user := model.User{
		ID:          ulid.Make().String(),
		Name:        data.Name,
		Username:    data.Username,
		Email:       data.Email,
		Password:    string(password),
		PhoneNumber: data.PhoneNumber,
		RoleID:      data.RoleID,
		Status:      data.Status,
		CreatedAt:   time.Now(),
		CreatedBy:   data.CreatedBy,
	}
	err := s.userRepo.Create(ctx, &user)
	if err != nil {
		return nil, tolo.NewError(http.StatusInternalServerError, "failed to create use", nil)
	}

	return &user, nil
}

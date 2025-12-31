package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/sugandasu/go-boilerplate/internal/app/model"
	"github.com/sugandasu/ruru/jongi"
	"github.com/sugandasu/ruru/tolo"
	"gorm.io/gorm"
)

func (s *service) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, tolo.NewError(http.StatusInternalServerError, "failed to get user", nil)
	}
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, tolo.NewError(http.StatusNotFound, tolo.NOT_FOUND, nil)
	}
	auth := jongi.GetAuthFromContext(ctx)
	if auth == nil {
		return nil, tolo.NewError(http.StatusBadRequest, "user is not found", nil)
	}

	return user, nil
}

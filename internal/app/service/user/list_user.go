package user

import (
	"context"

	"github.com/sugandasu/go-boilerplate/internal/app/model"
	"github.com/sugandasu/ruru/tolo"
)

func (s *service) ListUser(ctx context.Context, filters model.UserFilter) (tolo.PaginationResponse, error) {
	users, total, err := s.userRepo.Filter(ctx, filters)
	if err != nil {
		return tolo.NewPaginationResponse([]model.User{}, int(total)), tolo.NewError(500, "failed to list users", nil)
	}

	return tolo.NewPaginationResponse(users, int(total)), nil
}

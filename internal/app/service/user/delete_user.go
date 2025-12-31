package user

import (
	"context"
	"net/http"

	"github.com/sugandasu/ruru/tolo"
)

func (s *service) DeleteUser(ctx context.Context, id string) error {
	_, err := s.GetUserByID(ctx, id)
	if err != nil {
		return err
	}

	if err := s.userRepo.Delete(ctx, id); err != nil {
		return tolo.NewError(http.StatusInternalServerError, "failed to delete user", nil)
	}

	return nil
}

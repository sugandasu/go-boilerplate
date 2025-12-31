package user

import (
	"context"
	"net/http"

	"github.com/sugandasu/go-boilerplate/internal/app/model"
	"github.com/sugandasu/ruru/tolo"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) UpdateUser(ctx context.Context, data model.UserUpdateRequest) error {
	user, err := s.GetUserByID(ctx, data.ID)
	if err != nil {
		return err
	}

	if data.RoleID != nil {
		_, err = s.userRepo.GetRoleByID(ctx, *data.RoleID)
		if err != nil {
			return tolo.NewError(http.StatusBadRequest, "fail to get role", map[string]string{
				"role_id": "role is not found",
			})
		}
	}

	// Validation
	email, err := s.userRepo.Find(ctx, &model.User{Email: data.Email})
	if err != nil {
		return tolo.NewError(http.StatusInternalServerError, "failed to get user", nil)
	}
	if email.ID != user.ID {
		return tolo.NewError(http.StatusBadRequest, tolo.FAILED_VALIDATION, map[string]string{"email": "email must be unique"})
	}

	phoneNumber, err := s.userRepo.Find(ctx, &model.User{PhoneNumber: data.PhoneNumber})
	if err != nil {
		return tolo.NewError(http.StatusInternalServerError, "failed to get user", nil)
	}
	if phoneNumber.ID != user.ID {
		return tolo.NewError(http.StatusBadRequest, tolo.FAILED_VALIDATION, map[string]string{"phone_number": "phone_number must be unique"})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 0)

	user.Name = data.Name
	user.Username = data.Username
	user.Email = data.Email
	user.Password = string(password)
	user.PhoneNumber = data.PhoneNumber
	user.RoleID = data.RoleID
	user.Status = data.Status
	user.UpdatedAt = &data.UpdatedAt
	user.UpdatedBy = &data.UpdatedBy

	err = s.userRepo.Update(ctx, user)
	if err != nil {
		return tolo.NewError(http.StatusInternalServerError, "failed to update user", nil)
	}
	return nil
}

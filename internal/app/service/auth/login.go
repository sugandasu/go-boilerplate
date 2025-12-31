package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/oklog/ulid/v2"
	"github.com/sugandasu/go-boilerplate/internal/app/model"
	"github.com/sugandasu/ruru/jongi"
	"github.com/sugandasu/ruru/tolo"
)

func (s *service) Login(ctx context.Context, data model.LoginRequest) (*model.LoginResponse, error) {
	user, err := s.userRepository.Find(ctx, &model.User{
		Username: data.Username,
	})
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, tolo.NewError(http.StatusUnprocessableEntity, "invalid username or password", "user not found")
	}

	if !jongi.CheckPasswordHash(data.Password, user.Password) {
		return nil, tolo.NewError(http.StatusUnprocessableEntity, "invalid username or password", nil)
	}

	authRole, err := s.userRepository.GetRoleByID(ctx, *user.RoleID)
	if err != nil {
		return nil, tolo.NewError(http.StatusInternalServerError, "failed to get user roles", nil)
	}

	accessClaims := jongi.AuthClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:       ulid.Make().String(),
			Issuer:   s.config.App.Name,
			Audience: jwt.ClaimStrings{user.ID, "", data.Username},
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(s.config.Jwt.AccessTokenDuration),
			},
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
			NotBefore: &jwt.NumericDate{
				Time: time.Now(),
			},
		},
		UserID:    user.ID,
		CompanyID: "",
		Role:      *authRole,
	}
	accessToken, err := jongi.GenerateToken(accessClaims, s.config.Jwt.SecretKey)
	if err != nil {
		return nil, tolo.NewError(http.StatusInternalServerError, "failed to generate access token", err.Error())
	}

	refreshClaims := jongi.AuthClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:       ulid.Make().String(),
			Issuer:   s.config.App.Name,
			Subject:  "",
			Audience: jwt.ClaimStrings{user.ID, "", data.Username},
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(s.config.Jwt.RefreshTokenDuration),
			},
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
			NotBefore: &jwt.NumericDate{
				Time: time.Now(),
			},
		},
		UserID:    user.ID,
		CompanyID: "",
		Role:      *authRole,
	}
	refreshToken, err := jongi.GenerateToken(refreshClaims, s.config.Jwt.SecretKey)
	if err != nil {
		return nil, tolo.NewError(http.StatusInternalServerError, "failed to generate refresh token", err.Error())
	}

	res := &model.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	// TODO: Store the refresh token in the database or cache (e.g., Redis) for future validation and revocation.

	return res, nil
}

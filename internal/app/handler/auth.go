package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sugandasu/go-boilerplate/internal/app/model"
	"github.com/sugandasu/go-boilerplate/internal/app/service/auth"
	"github.com/sugandasu/ruru/tolo"
)

type AuthHandler interface {
	Login(c echo.Context) error
	RegisterRoutes(e *echo.Group)
}

type authHandler struct {
	authService auth.AuthService
}

func NewAuthHandler(authService auth.AuthService) AuthHandler {
	if authService == nil {
		panic("authService is nil")
	}

	return &authHandler{
		authService: authService,
	}
}

func (h *authHandler) RegisterRoutes(e *echo.Group) {
	e.POST("/auth/login", h.Login)
}

// Login godoc
// @Summary      Login
// @Description  Authenticate user and return token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body      model.LoginRequest  true  "Login credentials"
// @Success      200   {object}  dto.Response{Data=model.LoginResponse}
// @Failure      400   {object}  dto.Response
// @Failure      401   {object}  dto.Response
// @Router       /auth/login [post]
func (h *authHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()

	var req model.LoginRequest

	if err := c.Bind(&req); err != nil {
		tolo.ResponseError(c.Response().Writer, err)
		return err
	}

	err := tolo.Validator().Struct(&req)
	if err != nil {
		tolo.ResponseError(c.Response().Writer, tolo.NewError(http.StatusBadRequest, "failed to validate data", tolo.ValidatorTranslate(req, err)))
		return nil
	}

	res, err := h.authService.Login(ctx, req)
	if err != nil {
		return err
	}

	tolo.ResponseSuccess(c.Response().Writer, "login success", res)
	return nil
}

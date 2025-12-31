package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sugandasu/go-boilerplate/internal/app/model"
	"github.com/sugandasu/go-boilerplate/internal/app/service/user"
	"github.com/sugandasu/ruru/jongi"
	"github.com/sugandasu/ruru/tolo"
)

type UserHandler interface {
	RegisterRoutes(e *echo.Group)
	List(c echo.Context) error
	Get(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type userHandler struct {
	userService user.UserService
}

func NewUserHandler(userService user.UserService) UserHandler {
	if userService == nil {
		panic("userService is nil")
	}

	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) RegisterRoutes(e *echo.Group) {
	e.POST("/users", h.Create)
	e.GET("/users", h.List)
	e.GET("/users/:id", h.Get)
	e.PUT("/users/:id", h.Update)
	e.DELETE("/users/:id", h.Delete)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param user body model.UserCreateRequest true "User Create Request"
// @Success 200 {object} dto.Response{data=model.User} "Created"
// @Failure 400 {object} dto.Response{message=string} "Bad Request"
// @Failure 500 {object} dto.Response{message=string} "Internal Server Error"
// @Router /users [post]
func (h *userHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	var req model.UserCreateRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	auth := jongi.GetAuthFromContext(ctx)
	if auth == nil {
		return tolo.NewError(http.StatusBadRequest, "user is not found", nil)
	}
	req.CreatedBy = auth.UserID

	err := tolo.Validator().Struct(&req)
	if err != nil {
		return tolo.NewError(http.StatusBadRequest, "failed to validate data", tolo.ValidatorTranslate(req, err))
	}

	res, err := h.userService.CreateUser(ctx, req)
	if err != nil {
		return err
	}

	tolo.ResponseSuccess(c.Response().Writer, "success", res)
	return nil
}

// ListUser godoc
// @Summary List users
// @Description List user with filters
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param name query string false "name"
// @Param username query string false "username"
// @Param email query string false "email"
// @Param status query string false "status"
// @Param page query int false "page"
// @Param per_page query int false "per_page"
// @Success 200 {object} dto.Response{data=dto.PaginationResponse{Items=[]model.User}} "Success"
// @Failure 400 {object} dto.Response{message=string} "Bad Request"
// @Failure 500 {object} dto.Response{message=string} "Internal Server Error"
// @Router /users [get]
func (h *userHandler) List(c echo.Context) error {
	ctx := c.Request().Context()

	var req model.UserFilter
	if err := c.Bind(&req); err != nil {
		return err
	}

	auth := jongi.GetAuthFromContext(ctx)
	if auth == nil {
		return tolo.NewError(http.StatusBadRequest, "user is not found", nil)
	}

	if err := tolo.Validator().Struct(&req); err != nil {
		return err
	}

	res, err := h.userService.ListUser(ctx, req)
	if err != nil {
		return err
	}

	tolo.ResponseSuccess(c.Response().Writer, "success", res)
	return nil
}

// GerUser godoc
// @Summary Get a user
// @Description Get a user with id
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "User ID"
// @Success 200 {object} dto.Response{Data=model.User} "Success"
// @Failure 400 {object} dto.Response{message=string} "Bad Request"
// @Failure 500 {object} dto.Response{message=string} "Internal Server Error"
func (h *userHandler) Get(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if id == "" {
		return tolo.NewError(http.StatusBadRequest, "ID is required", nil)
	}

	user, err := h.userService.GetUserByID(ctx, id)
	if err != nil {
		return err
	}

	tolo.ResponseSuccess(c.Response().Writer, "success", user)
	return nil
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user with the provided information
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param user body model.UserUpdateRequest true "User Update Request"
// @Success 200 {object} dto.Response "Updated"
// @Failure 400 {object} dto.Response{message=string} "Bad Request"
// @Failure 500 {object} dto.Response{message=string} "Internal Server Error"
// @Router /users/{id} [put]
func (h *userHandler) Update(c echo.Context) error {
	ctx := c.Request().Context()

	var req model.UserUpdateRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	auth := jongi.GetAuthFromContext(ctx)
	if auth == nil {
		return tolo.NewError(http.StatusBadRequest, "user is not found", nil)
	}
	req.UpdatedBy = auth.UserID
	req.UpdatedAt = time.Now()

	if err := tolo.Validator().Struct(&req); err != nil {
		return err
	}

	err := h.userService.UpdateUser(ctx, req)
	if err != nil {
		return err
	}

	tolo.ResponseSuccess(c.Response().Writer, "success", nil)
	return nil
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by its ID
// @Tags Companies
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "User ID"
// @Success 200 {object} dto.Response{data=string} "Success"
// @Failure 400 {object} dto.Response{message=string} "Bad Request"
// @Failure 500 {object} dto.Response{message=string} "Internal Server Error"
// @Router /users/{id} [delete]
func (h *userHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if id == "" {
		return tolo.NewError(http.StatusBadRequest, "ID is required", nil)
	}

	err := h.userService.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	tolo.ResponseSuccess(c.Response().Writer, "success", nil)
	return nil
}

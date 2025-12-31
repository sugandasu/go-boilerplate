package model

import "time"

type User struct {
	ID          string     `gorm:"column:id" json:"id"`
	RoleID      *string    `gorm:"column:role_id" json:"role_id"`
	Name        string     `gorm:"column:name" json:"name"`
	Username    string     `gorm:"column:username" json:"username"`
	Email       string     `gorm:"column:email" json:"email"`
	Password    string     `gorm:"column:password" json:"password"`
	PhoneNumber string     `gorm:"column:phone_number" json:"phone_number"`
	Status      string     `gorm:"column:status" json:"status"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"created_at"`
	CreatedBy   string     `gorm:"column:created_by" json:"created_by"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy   *string    `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy   *string    `gorm:"column:deleted_by" json:"deleted_by"`
}

type UserCreateRequest struct {
	Name        string  `json:"name" validate:"required"`
	Username    string  `json:"username" validate:"required"`
	Email       string  `json:"email" validate:"required"`
	Password    string  `json:"password" validate:"required"`
	PhoneNumber string  `json:"phone_number" validate:"required"`
	RoleID      *string `json:"role_id"`
	Status      string  `json:"status" validate:"required"`
	CreatedBy   string  `validate:"required" swaggerignore:"true"`
}

type UserUpdateRequest struct {
	ID          string    `param:"id" validate:"required,ulid" swaggerignore:"true"`
	Name        string    `json:"name" validate:"required"`
	Gender      string    `json:"gender" validate:"required"`
	Username    string    `json:"username" validate:"required"`
	Email       string    `json:"email" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	PhoneNumber string    `json:"phone_number" validate:"required"`
	RoleID      *string   `json:"role_id"`
	Status      string    `json:"status" validate:"required"`
	UpdatedAt   time.Time `validate:"required" swaggerignore:"true"`
	UpdatedBy   string    `validate:"required" swaggerignore:"true"`
}

type UserFilter struct {
	Name     string `query:"name"`
	Username string `query:"username"`
	Email    string `query:"email"`
	Role     string `query:"role"`
	Status   string `query:"status"`
	Page     int    `query:"page"`
	PerPage  int    `query:"per_page"`
}

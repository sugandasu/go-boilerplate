package model

import "time"

type Role struct {
	ID        string     `gorm:"column:id"`
	Name      string     `gorm:"column:name"`
	Level     int        `gorm:"column:level"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	CreatedBy string     `gorm:"column:created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	UpdatedBy *string    `gorm:"column:updated_by"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
	DeletedBy *string    `gorm:"column:deleted_by"`
}

func (Role) TableName() string {
	return "roles"
}

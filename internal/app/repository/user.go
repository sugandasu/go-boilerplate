package repository

import (
	"context"
	"errors"

	"github.com/sugandasu/go-boilerplate/internal/app/model"
	"github.com/sugandasu/ruru/jongi"
	"github.com/sugandasu/ruru/nibirudb"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
	GetRoleByID(ctx context.Context, roleID string) (*jongi.AuthRole, error)
	GetByID(ctx context.Context, id string) (*model.User, error)
	Find(ctx context.Context, data *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Filter(ctx context.Context, data model.UserFilter) ([]model.User, int64, error)
}

type userRepository struct {
	db nibirudb.Database
}

func NewUserRepository(db nibirudb.Database) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.DB(ctx).Create(user).Error
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	if err := r.db.DB(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Find(ctx context.Context, data *model.User) (*model.User, error) {
	var user model.User
	if err := r.db.DB(ctx).Where(data).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	return r.db.DB(ctx).Save(user).Error
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	return r.db.DB(ctx).Where("id = ?", id).Delete(&model.User{}).Error
}

func (r *userRepository) GetRoleByID(ctx context.Context, roleID string) (*jongi.AuthRole, error) {
	var role model.Role
	err := r.db.DB(ctx).
		Table(model.Role{}.TableName()).
		Where("id = ?", roleID).
		First(&role).Error
	if err != nil {
		return nil, err
	}

	return &jongi.AuthRole{ID: role.ID, Name: role.Name, Level: role.Level}, nil
}

func (r *userRepository) Filter(ctx context.Context, data model.UserFilter) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	query := r.db.DB(ctx).Model(&model.User{})

	if data.Name != "" {
		query = query.Where("name ILIKE ?", "%"+data.Name+"%")
	}
	if data.Username != "" {
		query = query.Where("username ILIKE ?", "%"+data.Username+"%")
	}
	if data.Email != "" {
		query = query.Where("email ILIKE ?", "%"+data.Email+"%")
	}
	if data.Status != "" {
		query = query.Where("status = ?", data.Status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if data.Page == 0 {
		data.Page = 1
	}
	if data.PerPage == 0 {
		data.PerPage = 10
	}
	offset := (data.Page - 1) * data.PerPage
	query = query.Offset(offset).Limit(data.PerPage)

	err = query.Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

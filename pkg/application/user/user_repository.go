package user

import (
	domain "go-crud/pkg/domain/user"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetByUsername(username string) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByUsername(username string) (*domain.User, error) {
	var foundUser domain.User
	err := r.db.Where("username = ?", username).First(&foundUser).Error
	if err != nil {
		return nil, err
	}
	return &foundUser, nil
}

package repository

import (
	"uchiiParfume/features/users/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) entity.UsersRepositoryInterface {
	return &userRepository{
		db: db,
	}
}

// CreateUser implements entity.UsersRepositoryInterface.
func (u *userRepository) CreateUser(data entity.UsersCore) (entity.UsersCore, error) {
	panic("unimplemented")
}

// DeleteUser implements entity.UsersRepositoryInterface.
func (u *userRepository) DeleteUser(id string) error {
	panic("unimplemented")
}

// GetAllUser implements entity.UsersRepositoryInterface.
func (u *userRepository) GetAllUser() ([]entity.UsersCore, error) {
	panic("unimplemented")
}

// GetById implements entity.UsersRepositoryInterface.
func (u *userRepository) GetById(id string) (entity.UsersCore, error) {
	panic("unimplemented")
}

// UpdateUser implements entity.UsersRepositoryInterface.
func (u *userRepository) UpdateUser(id string, data entity.UsersCore) error {
	panic("unimplemented")
}



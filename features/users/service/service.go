package service

import "uchiiParfume/features/users/entity"

type userService struct {
	UserRepository entity.UsersRepositoryInterface
}

// CreateUser implements entity.UsersServiceInterface.
func (u *userService) CreateUser(data entity.UsersCore) (entity.UsersCore, error) {
	panic("unimplemented")
}

// DeleteUser implements entity.UsersServiceInterface.
func (u *userService) DeleteUser(id string) error {
	panic("unimplemented")
}

// GetAllUser implements entity.UsersServiceInterface.
func (u *userService) GetAllUser() ([]entity.UsersCore, error) {
	panic("unimplemented")
}

// GetById implements entity.UsersServiceInterface.
func (u *userService) GetById(id string) (entity.UsersCore, error) {
	panic("unimplemented")
}

// Login implements entity.UsersServiceInterface.
func (u *userService) Login(email string, password string) (entity.UsersCore, string, error) {
	panic("unimplemented")
}

// UpdateUser implements entity.UsersServiceInterface.
func (u *userService) UpdateUser(id string, data entity.UsersCore) error {
	panic("unimplemented")
}

func NewUserService(user entity.UsersRepositoryInterface) entity.UsersServiceInterface {
	return &userService{
		UserRepository: user,
	}
}

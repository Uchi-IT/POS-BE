package service

import (
	"errors"
	"regexp"
	"uchiiParfume/features/users/entity"
	crypt "uchiiParfume/utils/bcrypt"
)

type userService struct {
	UserRepository entity.UsersRepositoryInterface
}

func NewUserService(user entity.UsersRepositoryInterface) entity.UsersServiceInterface {
	return &userService{
		UserRepository: user,
	}
}

// CreateUser implements entity.UsersServiceInterface.
func (userUC *userService) CreateUser(data entity.UsersCore) (entity.UsersCore, error) {
	if data.Email == "" || data.Password == "" {
		return entity.UsersCore{}, errors.New("error, email or password can't be empty")
	}
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, data.Email)
	if !match {
		return entity.UsersCore{}, errors.New("error. email format not valid")
	}

	errRegister, err := userUC.UserRepository.CreateUser(data)
	if err != nil {
		return entity.UsersCore{}, err
	}

	return errRegister, nil
}

// DeleteUser implements entity.UsersServiceInterface.
func (userUC *userService) DeleteUser(id string) error {
	if id == "" {
		return errors.New("insert user id")
	}

	userData, err := userUC.UserRepository.GetById(id)
	if err != nil {
		return errors.New("user not found")
	}

	if userData.Id == "" {
		return errors.New("user not found")
	}

	errDelete := userUC.UserRepository.DeleteUser(id)
	if errDelete != nil {
		return errors.New("can't delete user")
	}

	return nil
}

// GetAllUser implements entity.UsersServiceInterface.
func (userUC *userService) GetAllUser() ([]entity.UsersCore, error) {
	users, err := userUC.UserRepository.GetAllUser()
	if err != nil {
		return nil, errors.New("error get data")
	}

	return users, nil
}

// GetById implements entity.UsersServiceInterface.
func (userUC *userService) GetById(id string) (entity.UsersCore, error) {
	if id == "" {
		return entity.UsersCore{}, errors.New("user ID is required")
	}

	user, err := userUC.UserRepository.GetById(id)
	if err != nil {
		return entity.UsersCore{}, err
	}

	return user, nil
}

// Login implements entity.UsersServiceInterface.
func (userUC *userService) Login(email string, password string) (entity.UsersCore, string, error) {
	if email == "" || password == "" {
		return entity.UsersCore{}, "", errors.New("error, email or password can't be empty")
	}

	loginData, token, err := userUC.UserRepository.Login(email, password)
	if err != nil {
		return entity.UsersCore{}, "", err
	}

	if crypt.CompareHash(loginData.Password, password) {
		return loginData, token, nil
	}

	return entity.UsersCore{}, "", errors.New("Login Failed")
}

// UpdateUser implements entity.UsersServiceInterface.
func (userUC *userService) UpdateUser(id string, data entity.UsersCore) error {
	if id == "" {
		return errors.New("id is not found")
	}

	user, errGet := userUC.UserRepository.GetById(id)
	if errGet != nil {
		return errGet
	}

	if user.Id == ""{
		return errors.New("user not found")
	}

	err := userUC.UserRepository.UpdateUser(id, data)
	if err != nil {
		return err
	}

	return nil

}

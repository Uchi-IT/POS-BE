package repository

import (
	"errors"
	"uchiiParfume/features/users/entity"
	"uchiiParfume/features/users/model"
	bcrypt "uchiiParfume/utils/bcrypt"
	utils "uchiiParfume/utils/jwt"

	"github.com/google/uuid"
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

// Login implements entity.UsersRepositoryInterface.
func (userRepo *userRepository) Login(email string, password string) (entity.UsersCore, string, error) {
	var data model.User

	bcrypt.CompareHash(data.Password, password)

	tx := userRepo.db.Where("email=?", email).First(&data)
	if tx.Error != nil {
		return entity.UsersCore{}, "", tx.Error
	}

	var token string

	if tx.RowsAffected > 0 {
		var errToken error
		token, errToken = utils.CreateToken(data.Id, data.Role)
		if errToken != nil {
			return entity.UsersCore{}, "", errToken
		}
	}

	var resp = entity.UserModelToUserCore(data)

	return resp, token, nil
}

// CreateUser implements entity.UsersRepositoryInterface.
func (userRepo *userRepository) CreateUser(data entity.UsersCore) (entity.UsersCore, error) {
	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return entity.UsersCore{}, UUIDerr
	}

	hashPassword, err := bcrypt.HashPassword(data.Password)
	if err != nil {
		return entity.UsersCore{}, err
	}

	var input = model.User{
		Id:       newUUID.String(),
		Email:    data.Email,
		Password: hashPassword,
		Role:     "user",
		Cabang:   data.Cabang,
	}

	erruser := userRepo.db.Save(&input)
	if erruser.Error != nil {
		return entity.UsersCore{}, erruser.Error
	}

	var resp = entity.UserModelToUserCore(input)

	return resp, nil
}

// DeleteUser implements entity.UsersRepositoryInterface.
func (userRepo *userRepository) DeleteUser(id string) error {
	var chekcId model.User

	errData := userRepo.db.Where("id = ?", id).Delete(&chekcId)
	if errData != nil {
		return errData.Error
	}

	return nil
}

// GetAllUser implements entity.UsersRepositoryInterface.
func (userRepo *userRepository) GetAllUser() ([]entity.UsersCore, error) {
	var dataUser []model.User

	errData := userRepo.db.Find(&dataUser).Error
	if errData != nil {
		return nil, errData
	}

	mapData := entity.ListUserModelToUserCore(dataUser)
	return mapData, nil
}

// GetById implements entity.UsersRepositoryInterface.
func (userRepo *userRepository) GetById(id string) (entity.UsersCore, error) {
	var data model.User
	errData := userRepo.db.Where("id=?", id).First(&data).Error
	if errData != nil {
		return entity.UsersCore{}, errData
	}

	resp := entity.UserModelToUserCore(data)

	return resp, nil
}

// UpdateUser implements entity.UsersRepositoryInterface.
func (userRepo *userRepository) UpdateUser(id string, data entity.UsersCore) error {

	request := entity.UserCoreToUserModel(data)

	tx := userRepo.db.Where("id = ?", id).Updates(&request)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

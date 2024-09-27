package repository

import (
	"errors"
	"fmt"
	cabang "uchiiParfume/features/cabang/entity"
	"uchiiParfume/features/users/entity"
	"uchiiParfume/features/users/model"
	bcrypt "uchiiParfume/utils/bcrypt"
	utils "uchiiParfume/utils/jwt"
	"uchiiParfume/utils/pagination"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db     *gorm.DB
	cabang cabang.CabangRepositoryInterface
}

func NewUserRepository(db *gorm.DB, cabang cabang.CabangRepositoryInterface) entity.UsersRepositoryInterface {
	return &userRepository{
		db:     db,
		cabang: cabang,
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
	userData := entity.UserCoreToUserModel(data)
	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return entity.UsersCore{}, UUIDerr
	}

	hashPassword, errHash := bcrypt.HashPassword(data.Password)
	if errHash != nil {
		return entity.UsersCore{}, errHash
	}

	userData.Id = newUUID.String()
	userData.Password = hashPassword

	txOuter := userRepo.db.Begin()

	if err := txOuter.Save(&userData).Error; err != nil {
		txOuter.Rollback()
		return entity.UsersCore{}, err
	}

	input := entity.UserModelToUserCore(userData)
	input.Role = "user"

	for i, cabangId := range data.Cabang_id {
		_, tx := userRepo.cabang.GetById(cabangId)
		if tx != nil {
			txOuter.Rollback()
			return entity.UsersCore{}, errors.New("cabang tidak ada")
		}

		dataCabang := new(model.UserCabang)
		dataCabang.UserID = userData.Id
		dataCabang.CabangID = cabangId

		fmt.Printf("Saving UserCabang: UserID=%s, CabangID=%s\n", dataCabang.UserID, dataCabang.CabangID)

		for j := i + 1; j < len(data.Cabang_id); j++ {
			if cabangId == data.Cabang_id[j] {
				return entity.UsersCore{}, errors.New("error : cabang tidak boleh sama")
			}
		}
		txInner := txOuter.Create(&dataCabang)
		if txInner.Error != nil {
			fmt.Printf("Error saving UserCabang: %v\n", txInner.Error)
			txOuter.Rollback()
			return entity.UsersCore{}, txInner.Error
		}

	}

	txOuter.Commit()

	var resp = entity.UsersCore{
		Id:     input.Id,
		Email:  input.Email,
		Role:   input.Role,
		Cabang: input.Cabang,
	}

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
func (userRepo *userRepository) GetAllUser(page, limit int, search, sort, filter string) ([]entity.UsersCore, pagination.PageInfo, int, error) {
	var dataUser []model.User
	var totalCount int64
	offset := (page - 1) * limit

	query := userRepo.db.Preload("Cabang").Model(&model.User{}).
		Joins("JOIN user_cabangs ON user_cabangs.user_id = users.id")

	if search != "" {
		query = query.Where("users.nama LIKE ? OR users.email LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if filter != "" {
		query = query.Where("user_cabangs.cabang_id = ?", filter)
	}

	if sort == "nama_asc" {
		query = query.Order("users.nama ASC")
	} else if sort == "nama_desc" {
		query = query.Order("users.nama DESC")
	} else {
		query = query.Order("users.created_at DESC")
	}

	tx := query.Count(&totalCount)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	tx = query.Offset(offset).Limit(limit).Find(&dataUser)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	mapData := entity.ListUserModelToUserCore(dataUser)

	pageInfo := pagination.CalculateData(int(totalCount), limit, page)

	return mapData, pageInfo, int(totalCount), nil
}

// GetById implements entity.UsersRepositoryInterface.
func (userRepo *userRepository) GetById(id string) (entity.UsersCore, error) {
	var data model.User
	errData := userRepo.db.Preload("Cabang").Where("id=?", id).First(&data).Error
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

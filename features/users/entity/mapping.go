package entity

import (
	"uchiiParfume/features/users/model"
)

func UserModelToUserCore(user model.User) UsersCore {
	userCore := UsersCore{
		Id:       user.Id,
		Email:    user.Email,
		Password: user.Password,
		RoleId:   user.RoleId,
		Cabang:   user.Cabang,
		Role:     user.Role,
	}
	return userCore
}

func UserCoreToUserModel(user UsersCore) model.User {
	userModel := model.User{
		Id:       user.Id,
		Email:    user.Email,
		Password: user.Password,
		RoleId:   user.RoleId,
		Cabang:   user.Cabang,
		Role:     user.Role,
	}
	return userModel
}

func ListUserModelToUserCore(user []model.User) []UsersCore {
	coreUser := []UsersCore{}
	for _, v := range user {
		users := UserModelToUserCore(v)
		coreUser = append(coreUser, users)
	}
	return coreUser
}

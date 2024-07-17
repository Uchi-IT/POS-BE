package entity

import (
	"uchiiParfume/features/users/model"
	cm "uchiiParfume/features/cabang/model"	
)

func CabangModelToCabangCore(data cm.Cabang) UserCabangCore{
	return UserCabangCore{
		Cabang: data.NamaCabang,
	}
}

func ListCabangModelToListCabangCore(data []cm.Cabang) []UserCabangCore{
	coreCabang := []UserCabangCore{}
	for _, v := range data {
		cabang := CabangModelToCabangCore(v)
		coreCabang = append(coreCabang, cabang)
	}
	return coreCabang
}

func CabangCoreToCabangModel(data UserCabangCore) cm.Cabang {
	return cm.Cabang{
		// ID:        category.TrashCategoryID,
		NamaCabang: data.Cabang,
	}
}

func ListCabangCoreToCabangModel(data []UserCabangCore) []cm.Cabang {
	coreCabang := []cm.Cabang{}
	for _, v := range data {
		categorys := CabangCoreToCabangModel(v)
		coreCabang = append(coreCabang, categorys)
	}
	return coreCabang
}

func UserModelToUserCore(user model.User) UsersCore {
	userCore := UsersCore{
		Id:       user.Id,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}
	cabang := ListCabangModelToListCabangCore(user.Cabang)
	userCore.Cabang = cabang
	return userCore
}

func UserCoreToUserModel(user UsersCore) model.User {
	userModel := model.User{
		Id:       user.Id,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}
	cabang := ListCabangCoreToCabangModel(user.Cabang)
	userModel.Cabang = cabang
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

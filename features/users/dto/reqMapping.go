package handler

import "uchiiParfume/features/users/entity"


func CabangrequestToCabangCore(data UserCabangRequest) entity.UserCabangCore {
	return entity.UserCabangCore{
		// TrashCategoryID: category.TrashCategoryID,
		Cabang: data.Cabang,
	}
}

func ListCabangRequestToCabangCore(data []UserCabangRequest) []entity.UserCabangCore {
	listCategory := []entity.UserCabangCore{}
	for _, v := range data {
		category := CabangrequestToCabangCore(v)
		listCategory = append(listCategory, category)
	}

	return listCategory
}

func UserRequestToUserCore(data UserRequest) entity.UsersCore {
	cabangReq := entity.UsersCore{
		Email:    data.Email,
		Password: data.Password,
		Cabang_id: data.Cabang_id,
	}
	cabang := ListCabangRequestToCabangCore(data.Cabang)
	cabangReq.Cabang = cabang
	return cabangReq
}
package handler

import "uchiiParfume/features/users/entity"

func CabangrequestToCabangCore(data UserCabangRequest) entity.UserCabangCore {
	return entity.UserCabangCore{
		// TrashCategoryID: category.TrashCategoryID,
		Cabang: data.Cabang,
	}
}

func ListCabangRequestToCabangCore(data []UserCabangRequest) []entity.UserCabangCore {
	listCabang := []entity.UserCabangCore{}
	for _, v := range data {
		cabang := CabangrequestToCabangCore(v)
		listCabang = append(listCabang, cabang)
	}

	return listCabang
}

func UserRequestToUserCore(data UserRequest) entity.UsersCore {
	cabangReq := entity.UsersCore{
		Nama:      data.Nama,
		Email:     data.Email,
		Password:  data.Password,
		Cabang_id: data.Cabang_id,
	}
	cabang := ListCabangRequestToCabangCore(data.Cabang)
	cabangReq.Cabang = cabang
	return cabangReq
}

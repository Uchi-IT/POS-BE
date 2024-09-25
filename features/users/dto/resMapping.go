package handler

import "uchiiParfume/features/users/entity"

func CabangCoreToCabangResponse(data entity.UserCabangCore) CabangResponse {
	return CabangResponse{
		Cabang: data.Cabang,
	}
}

func ListCabangCoreToCabangResponse(data []entity.UserCabangCore) []CabangResponse {
	ResponseCabang := []CabangResponse{}
	for _, v := range data {
		cabang := CabangCoreToCabangResponse(v)
		ResponseCabang = append(ResponseCabang, cabang)
	}
	return ResponseCabang
}

func UserCoreToUserResponse(data entity.UsersCore) UserResponse {
	userResp := UserResponse{
		Id:        data.Id,
		Nama:      data.Nama,
		Email:     data.Email,
		Role:      data.Role,
		Cabang_id: data.Cabang_id,
	}
	cabang := ListCabangCoreToCabangResponse(data.Cabang)
	userResp.Cabang = cabang

	return userResp
}

func ListUserCoreToListUserResponse(data []entity.UsersCore) []UserResponse {
	userResp := []UserResponse{}
	for _, user := range data {
		dataUser := UserCoreToUserResponse(user)
		userResp = append(userResp, dataUser)
	}
	return userResp
}

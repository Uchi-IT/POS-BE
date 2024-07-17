package handler

import "uchiiParfume/features/users/entity"

func CabangCoreToCabangResponse(data entity.UserCabangCore) CabangResponse {
	return CabangResponse{
		Cabang: data.Cabang,
	}
}

func ListCabangCoreToCabangResponse(data []entity.UserCabangCore) []CabangResponse {
	ResponseCategory := []CabangResponse{}
	for _, v := range data {
		category := CabangCoreToCabangResponse(v)
		ResponseCategory = append(ResponseCategory, category)
	}
	return ResponseCategory
}

func UserCoreToUserResponse(data entity.UsersCore) UserResponse {
	userResp := UserResponse{
		Id:    data.Id,
		Email: data.Email,
		Role:  data.Role,
		Cabang_id: data.Cabang_id,
	}
	cabang := ListCabangCoreToCabangResponse(data.Cabang)
	userResp.Cabang = cabang

	return userResp
}

func ListUserCoreToListUserResponse(data []entity.UsersCore) []UserResponse {
	userResp := []UserResponse{}
	for _, user := range data {
		dataData := UserCoreToUserResponse(user)
		userResp = append(userResp, dataData)
	}
	return userResp
}

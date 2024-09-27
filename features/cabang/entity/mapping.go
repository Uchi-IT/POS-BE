package entity

import "uchiiParfume/features/cabang/model"

func CabangModelToCabangCore(cabang model.Cabang) CabangCore {
	cabangCore := CabangCore{
		Id:         cabang.Id,
		Image:      cabang.Image,
		NamaCabang: cabang.NamaCabang,
		Alamat:     cabang.Alamat,
		CreatedAt:  cabang.CreatedAt,
		UpdatedAt:  cabang.UpdatedAt,
		DeleteAt:   cabang.DeleteAt,
	}
	return cabangCore
}

func CabangCoreToCabangModel(cabang CabangCore) model.Cabang {
	cabangModel := model.Cabang{
		Id:         cabang.Id,
		Image:      cabang.Image,
		NamaCabang: cabang.NamaCabang,
		Alamat:     cabang.Alamat,
		CreatedAt:  cabang.CreatedAt,
		UpdatedAt:  cabang.UpdatedAt,
		DeleteAt:   cabang.DeleteAt,
	}
	return cabangModel
}

func ListCabangModelToCabangCore(cabang []model.Cabang) []CabangCore {
	coreCabang := []CabangCore{}
	for _, v := range cabang {
		branchs := CabangModelToCabangCore(v)
		coreCabang = append(coreCabang, branchs)
	}
	return coreCabang
}

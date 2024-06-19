package entity

import "uchiiParfume/features/cabang/model"

func CabangModelToCabangCore(cabang model.Cabang) CabangCore {
	cabangCore := CabangCore{
		Id:         cabang.Id,
		NamaCabang: cabang.NamaCabang,
		Alamat:     cabang.Alamat,
	}
	return cabangCore
}

func CabangCoreToCabangModel(cabang CabangCore) model.Cabang {
	cabangModel := model.Cabang{
		Id:         cabang.Id,
		NamaCabang: cabang.NamaCabang,
		Alamat:     cabang.Alamat,
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

package entity

import "uchiiParfume/features/produkCabang/model"

func ProdukCabangModelToProdukCabangCore(produk model.ProdukCabang) ProdukCabangCore {
	produkCore := ProdukCabangCore{
		Id:                  produk.Id,
		Foto:                produk.Foto,
		NamaProduk:          produk.NamaProduk,
		HargaJual:           produk.HargaJual,
		SeratusMl:           produk.SeratusMl,
		DuaratusMl:          produk.DuaratusMl,
		DuaRatusLimaPuluhMl: produk.DuaRatusLimaPuluhMl,
		Manual:              produk.Manual,
		Total:               produk.Total,
	}
	return produkCore
}

func ProdukCabangCoreToProdukCabangModel(produk ProdukCabangCore) model.ProdukCabang {
	produkModel := model.ProdukCabang{
		Id:                  produk.Id,
		Foto:                produk.Foto,
		NamaProduk:          produk.NamaProduk,
		HargaJual:           produk.HargaJual,
		SeratusMl:           produk.SeratusMl,
		DuaratusMl:          produk.DuaratusMl,
		DuaRatusLimaPuluhMl: produk.DuaRatusLimaPuluhMl,
		Manual:              produk.Manual,
		Total:               produk.Total,
	}
	return produkModel
}

func ListProdukCabangModelToListProdukCabangCore(produk []model.ProdukCabang) []ProdukCabangCore {
	coreProduk := []ProdukCabangCore{}
	for _, v := range produk {
		data := ProdukCabangModelToProdukCabangCore(v)
		coreProduk = append(coreProduk, data)
	}
	return coreProduk
}

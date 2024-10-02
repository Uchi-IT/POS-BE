package entity

import (
	"uchiiParfume/features/produkCabang/model"
)

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
		CabangId:            produk.CabangId,
		ProdukId:            produk.ProdukId,
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
		CabangId:            produk.CabangId,
		ProdukId:            produk.ProdukId,
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

func ItemModelToItemCore(data model.RiwayatProdukCabangItem) RiwayatProdukCabangItemCore {
	return RiwayatProdukCabangItemCore{
		ProdukCabangId:      data.ProdukCabangId,
		NamaProduk:          data.NamaProduk,
		SeratusMl:           data.SeratusMl,
		DuaratusMl:          data.DuaratusMl,
		DuaRatusLimaPuluhMl: data.DuaRatusLimaPuluhMl,
		Manual:              data.Manual,
	}
}

func ListItemModelToItemCore(data []model.RiwayatProdukCabangItem) []RiwayatProdukCabangItemCore {
	coreItem := []RiwayatProdukCabangItemCore{}
	for _, v := range data {
		item := ItemModelToItemCore(v)
		coreItem = append(coreItem, item)
	}
	return coreItem
}

func ItemCoreToItemModel(data RiwayatProdukCabangItemCore) model.RiwayatProdukCabangItem {
	return model.RiwayatProdukCabangItem{
		ProdukCabangId:      data.ProdukCabangId,
		NamaProduk:          data.NamaProduk,
		SeratusMl:           data.SeratusMl,
		DuaratusMl:          data.DuaratusMl,
		DuaRatusLimaPuluhMl: data.DuaRatusLimaPuluhMl,
		Manual:              data.Manual,
	}
}

func ListItemCoreToItemModel(data []RiwayatProdukCabangItemCore) []model.RiwayatProdukCabangItem {
	coreItem := []model.RiwayatProdukCabangItem{}
	for _, v := range data {
		item := ItemCoreToItemModel(v)
		coreItem = append(coreItem, item)
	}
	return coreItem
}

func RiwayatModelToRiwayatCore(riwayat model.RiwayatProdukCabang) RiwayatProdukCabangCore {
	riwayatCore := RiwayatProdukCabangCore{
		Id:         riwayat.Id,
		NamaCabang: riwayat.NamaCabang,
		NamaAdmin:  riwayat.NamaAdmin,
		TotalStok:  riwayat.TotalStok,
		Catatan:    riwayat.Catatan,
		CreatedAt:  riwayat.CreatedAt,
		UpdatedAt:  riwayat.UpdatedAt,
	}
	item := ListItemModelToItemCore(riwayat.ProdukDetail)
	riwayatCore.ProdukDetail = item
	return riwayatCore
}

func RiwayatCoreToRiwayatModel(riwayat RiwayatProdukCabangCore) model.RiwayatProdukCabang {
	riwayatCore := model.RiwayatProdukCabang{
		Id:         riwayat.Id,
		NamaCabang: riwayat.NamaCabang,
		NamaAdmin:  riwayat.NamaAdmin,
		TotalStok:  riwayat.TotalStok,
		Catatan:    riwayat.Catatan,
		CreatedAt:  riwayat.CreatedAt,
		UpdatedAt:  riwayat.UpdatedAt,
	}
	item := ListItemCoreToItemModel(riwayat.ProdukDetail)
	riwayatCore.ProdukDetail = item
	return riwayatCore
}

func ListHistoryModelToHistoryCore(riwayat []model.RiwayatProdukCabang) []RiwayatProdukCabangCore {
	historyCore := []RiwayatProdukCabangCore{}
	for _, v := range riwayat {
		history := RiwayatModelToRiwayatCore(v)
		historyCore = append(historyCore, history)
	}
	return historyCore
}

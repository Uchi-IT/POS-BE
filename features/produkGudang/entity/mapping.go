package entity

import (
	"uchiiParfume/features/produkGudang/model"
)

func ProdukGModelToProdukGCore(produkG model.ProdukGudang) ProdukGudangCore {
	produkCore := ProdukGudangCore{
		Id:         produkG.Id,
		Foto:       produkG.Foto,
		NamaProduk: produkG.NamaProduk,
		Stok:       produkG.Stok,
		HargaJual:  produkG.HargaJual,
		HargaBeli:  produkG.HargaBeli,
	}
	return produkCore
}

func ProdukGCoreToProdukGModel(produkG ProdukGudangCore) model.ProdukGudang {
	produkModel := model.ProdukGudang{
		Id:         produkG.Id,
		Foto:       produkG.Foto,
		NamaProduk: produkG.NamaProduk,
		Stok:       produkG.Stok,
		HargaJual:  produkG.HargaJual,
		HargaBeli:  produkG.HargaBeli,
	}
	return produkModel
}

func ListProdukGModelToProdukGCore(produkG []model.ProdukGudang) []ProdukGudangCore {
	coreProduk := []ProdukGudangCore{}
	for _, v := range produkG {
		products := ProdukGModelToProdukGCore(v)
		coreProduk = append(coreProduk, products)
	}
	return coreProduk
}

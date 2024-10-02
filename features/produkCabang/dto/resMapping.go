package dto

import "uchiiParfume/features/produkCabang/entity"

func ItemCoreToItemResponse(data entity.RiwayatProdukCabangItemCore) ProdukCabangHistoryItemResponse {
	return ProdukCabangHistoryItemResponse{
		NamaProduk:          data.NamaProduk,
		SeratusMl:           data.SeratusMl,
		DuaratusMl:          data.DuaratusMl,
		DuaRatusLimaPuluhMl: data.DuaRatusLimaPuluhMl,
		Manual:              data.Manual,
	}
}

func ListItemCoreToItemResponse(data []entity.RiwayatProdukCabangItemCore) []ProdukCabangHistoryItemResponse {
	ResponseItem := []ProdukCabangHistoryItemResponse{}
	for _, v := range data {
		item := ItemCoreToItemResponse(v)
		ResponseItem = append(ResponseItem, item)
	}
	return ResponseItem
}

func RiwayatCoreToRiwayatResponse(data entity.RiwayatProdukCabangCore) ProdukCabangHistoryResponse {
	riwayatResp := ProdukCabangHistoryResponse{
		Id:        data.Id,
		NamaAdmin: data.NamaAdmin,
		TotalStok: data.TotalStok,
		Catatan:   data.Catatan,
		CreatedAt: data.CreatedAt,
	}
	parfumDetail := ListItemCoreToItemResponse(data.ProdukDetail)
	riwayatResp.Produk = parfumDetail

	return riwayatResp
}

func ListTransactionCoreToListTransactionResponse(data []entity.RiwayatProdukCabangCore) []ProdukCabangHistoryResponse {
	historyResp := []ProdukCabangHistoryResponse{}
	for _, history := range data {
		dataHistory := RiwayatCoreToRiwayatResponse(history)
		historyResp = append(historyResp, dataHistory)
	}
	return historyResp
}

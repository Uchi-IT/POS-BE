package dto

import "uchiiParfume/features/transaction/entity"

func ItemCoreToItemResponse(data entity.TransactionItemCore) ParfumItemResponse {
	return ParfumItemResponse{
		Nama:       data.Nama,
		Foto:       data.Foto,
		Harga:      data.Harga,
		Jumlah:     data.Jumlah,
		HargaTotal: data.HargaTotal,
	}
}

func ListItemCoreToItemResponse(data []entity.TransactionItemCore) []ParfumItemResponse {
	ResponseItem := []ParfumItemResponse{}
	for _, v := range data {
		item := ItemCoreToItemResponse(v)
		ResponseItem = append(ResponseItem, item)
	}
	return ResponseItem
}

func TransactionCoreToTransactionResponse(data entity.TransactionCore) TransactionResponse {
	transactionResp := TransactionResponse{
		Id:        data.Id,
		Nama:      data.Nama,
		Parfum_id: data.Parfum_id,
	}
	parfum := ListItemCoreToItemResponse(data.Parfum)
	transactionResp.Parfum = parfum

	return transactionResp
}

func ListTransactionCoreToListTransactionResponse(data []entity.TransactionCore) []TransactionResponse {
	transactionResp := []TransactionResponse{}
	for _, transaction := range data {
		dataTransaction := TransactionCoreToTransactionResponse(transaction)
		transactionResp = append(transactionResp, dataTransaction)
	}
	return transactionResp
}

func ListTransactionItemCoreToListTransactionItemResponse(data []entity.TransactionItemCore) []ParfumItemResponse {
	response := make([]ParfumItemResponse, len(data))
	for i, v := range data {
		response[i] = ParfumItemResponse{
			Nama:       v.Nama,
			Foto:       v.Foto,
			Harga:      v.Harga,
			Jumlah:     v.Jumlah,
			HargaTotal: v.HargaTotal,
		}
	}
	return response
}

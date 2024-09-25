package dto

import "uchiiParfume/features/transaction/entity"

func ItemRequestToItemCore(data ParfumItemRequest) entity.TransactionItemCore {
	return entity.TransactionItemCore{
		Nama:  data.Nama,
		Foto:  data.Foto,
		Harga: data.Harga,
	}
}

func ListItemRequestToListItemCore(data []ParfumItemRequest) []entity.TransactionItemCore {
	listItem := []entity.TransactionItemCore{}
	for _, v := range data {
		cabang := ItemRequestToItemCore(v)
		listItem = append(listItem, cabang)
	}

	return listItem
}

func TransactionRequestToTransactionCore(req TransactionRequest) entity.TransactionCore {
	items := []entity.TransactionItemCore{}
	for _, item := range req.Parfum {
		items = append(items, entity.TransactionItemCore{
			ProdukCabangId: item.ProdukCabangId,
			Jumlah:   item.Jumlah,
		})
	}
	return entity.TransactionCore{
		Parfum: items,
	}
}

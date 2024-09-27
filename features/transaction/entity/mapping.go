package entity

import (
	pc "uchiiParfume/features/produkCabang/model"
	"uchiiParfume/features/transaction/model"
)

func ItemModelToItemCore(data pc.ProdukCabang) TransactionItemCore {
	return TransactionItemCore{
		Nama:  data.NamaProduk,
		Foto:  data.Foto,
		Harga: data.HargaJual,
	}
}

// untuk nambahin jumlah dan harga total
func ItemDetailModelToItemDetailCore(data model.TransactionItem) TransactionItemCore {
	return TransactionItemCore{
		Nama:       data.Nama,
		Foto:       data.Foto,
		Harga:      data.Harga,
		Jumlah:     data.Jumlah,
		HargaTotal: data.HargaTotal,
	}
}

func ListItemModelToItemCore(data []pc.ProdukCabang) []TransactionItemCore {
	coreItem := []TransactionItemCore{}
	for _, v := range data {
		item := ItemModelToItemCore(v)
		coreItem = append(coreItem, item)
	}
	return coreItem
}

// untuk nambahin jumlah dan harga total
func ListItemDetailModelToItemDetailCore(data []model.TransactionItem) []TransactionItemCore {
	coreItem := []TransactionItemCore{}
	for _, v := range data {
		item := ItemDetailModelToItemDetailCore(v)
		coreItem = append(coreItem, item)
	}
	return coreItem
}

func ItemCoreToItemModel(data TransactionItemCore) pc.ProdukCabang {
	return pc.ProdukCabang{
		NamaProduk: data.Nama,
		Foto:       data.Foto,
		HargaJual:  data.Harga,
	}
}

func ItemDetailCoreToItemDetailModel(data TransactionItemCore) model.TransactionItem {
	return model.TransactionItem{
		Nama:       data.Nama,
		Foto:       data.Foto,
		Harga:      data.Harga,
		Jumlah:     data.Jumlah,
		HargaTotal: data.HargaTotal,
	}
}

func ListItemCoreToItemModel(data []TransactionItemCore) []pc.ProdukCabang {
	coreItem := []pc.ProdukCabang{}
	for _, v := range data {
		item := ItemCoreToItemModel(v)
		coreItem = append(coreItem, item)
	}
	return coreItem
}

func ListItemDetailCoreToItemDetailModel(data []TransactionItemCore) []model.TransactionItem {
	coreItem := []model.TransactionItem{}
	for _, v := range data {
		item := ItemDetailCoreToItemDetailModel(v)
		coreItem = append(coreItem, item)
	}
	return coreItem
}

func TransactionModelToTransactionCore(transaction model.Transaction) TransactionCore {
	transactionCore := TransactionCore{
		Id:         transaction.Id,
		Nama:       transaction.Nama,
		TotalHarga: transaction.TotalHarga,
	}
	item := ListItemModelToItemCore(transaction.Parfum)
	transactionCore.Parfum = item
	return transactionCore
}

func TransactionDetailModelToTransactionDetailCore(transaction model.Transaction) TransactionCore {
	transactionCore := TransactionCore{
		Id:         transaction.Id,
		Nama:       transaction.Nama,
		TotalHarga: transaction.TotalHarga,
	}
	item := ListItemDetailModelToItemDetailCore(transaction.ParfumDetail)
	transactionCore.ParfumDetail = item
	return transactionCore
}

func TransactionCoreToTransactionModel(transaction TransactionCore) model.Transaction {
	transactionModel := model.Transaction{
		Id:         transaction.Id,
		Nama:       transaction.Nama,
		TotalHarga: transaction.TotalHarga,
	}
	item := ListItemCoreToItemModel(transaction.Parfum)
	transactionModel.Parfum = item
	return transactionModel
}

func TransactionDetailCoreToTransactionDetailModel(transaction TransactionCore) model.Transaction {
	transactionModel := model.Transaction{
		Id:         transaction.Id,
		Nama:       transaction.Nama,
		TotalHarga: transaction.TotalHarga,
	}
	item := ListItemDetailCoreToItemDetailModel(transaction.ParfumDetail)
	transactionModel.ParfumDetail = item
	return transactionModel
}

func ListTransactionModelToTransactionCore(transaction []model.Transaction) []TransactionCore {
	transactionCore := []TransactionCore{}
	for _, v := range transaction {
		transactions := TransactionModelToTransactionCore(v)
		transactionCore = append(transactionCore, transactions)
	}
	return transactionCore
}

func ListTransactionDetailModelToTransactionDetailCore(transaction []model.Transaction) []TransactionCore {
	transactionCore := []TransactionCore{}
	for _, v := range transaction {
		transactions := TransactionDetailModelToTransactionDetailCore(v)
		transactionCore = append(transactionCore, transactions)
	}
	return transactionCore
}

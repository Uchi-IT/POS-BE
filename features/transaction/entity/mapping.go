package entity

import (
	"uchiiParfume/features/transaction/model"
	pc "uchiiParfume/features/produkCabang/model"
)

func ItemModelToItemCore(data pc.ProdukCabang) TransactionItemCore {
	return TransactionItemCore{
		Nama:  data.NamaProduk,
		Foto:  data.Foto,
		Harga: data.HargaJual,
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

func ItemCoreToItemModel(data TransactionItemCore) pc.ProdukCabang {
	return pc.ProdukCabang{
		NamaProduk: data.Nama,
		Foto:       data.Foto,
		HargaJual:  data.Harga,
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

func TransactionModelToTransactionCore(transaction model.Transaction) TransactionCore {
	transactionCore := TransactionCore{
		Id:   transaction.Id,
		Nama: transaction.Nama,
	}
	item := ListItemModelToItemCore(transaction.Parfum)
	transactionCore.Parfum = item
	return transactionCore
}

func TransactionCoreToTransactionModel(transaction TransactionCore) model.Transaction {
	transactionModel := model.Transaction{
		Id:   transaction.Id,
		Nama: transaction.Nama,
	}
	item := ListItemCoreToItemModel(transaction.Parfum)
	transactionModel.Parfum = item
	return transactionModel
}

func ListTransactionModelToTransactionCore(transaction []model.Transaction) []TransactionCore{
	transactionCore := []TransactionCore{}
	for _, v := range transaction {
		transactions := TransactionModelToTransactionCore(v)
		transactionCore = append(transactionCore, transactions)
	}
	return transactionCore
}
package dto

type TransactionRequest struct {
	Nama      string `json:"nama"`
	// Parfum_id []string `json:"parfum_id"`
	Parfum    []ParfumItemRequest  `json:"parfum"`
}

type ParfumItemRequest struct {
	ProdukCabangId string `json:"parfum_id"`
	Nama       string `json:"nama"`
	Foto       string `json:"foto"`
	Harga      int    `json:"harga"`
	Jumlah     int    `json:"jumlah"`
	HargaTotal int    `json:"harga_total"`
}
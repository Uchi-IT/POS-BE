package dto

import "time"

type ProdukCabangResponse struct {
	Id                  string `json:"id"`
	Foto                string `json:"foto"`
	NamaProduk          string `json:"nama_produk"`
	HargaJual           int    `json:"harga_jual"`
	SeratusMl           int    `json:"seratus_ml"`
	DuaratusMl          int    `json:"duaratus_ml"`
	DuaRatusLimaPuluhMl int    `json:"duaratuslimapuluh_ml"`
	Manual              int    `json:"manual"`
	Total               int    `json:"total"`
}

type ProdukCabangHistoryResponse struct {
	Id        uint      `json:"id"`
	NamaAdmin string    `json:"nama_admin"`
	TotalStok int       `json:"total_stok"`
	Catatan   string    `json:"catatan"`
	CreatedAt time.Time `json:"created_at"`
	Produk []ProdukCabangHistoryItemResponse
}

type ProdukCabangHistoryItemResponse struct {
	NamaProduk string `json:"nama_produk"`
	SeratusMl           int    `json:"seratus_ml"`
	DuaratusMl          int    `json:"duaratus_ml"`
	DuaRatusLimaPuluhMl int    `json:"duaratuslimapuluh_ml"`
	Manual              int    `json:"manual"`
}
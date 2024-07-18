package dto

type ProdukCabangRequest struct {
	Foto                string `json:"foto"`
	NamaProduk          string `json:"nama_produk"`
	HargaJual           int    `json:"harga_jual"`
	SeratusMl           int    `json:"seratus_ml"`
	DuaratusMl          int    `json:"duaratus_ml"`
	DuaRatusLimaPuluhMl int    `json:"duaratuslimapuluh_ml"`
	Manual              int    `json:"maual"`
}

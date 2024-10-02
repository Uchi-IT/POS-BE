package dto

type ProdukCabangRequest struct {
	Foto                string `json:"foto"`
	HargaJual           int    `json:"harga_jual"`
	SeratusMl           int    `json:"seratus_ml"`
	DuaratusMl          int    `json:"duaratus_ml"`
	DuaRatusLimaPuluhMl int    `json:"duaratuslimapuluh_ml"`
	Manual              int    `json:"manual"`
	ProdukId            string `json:"product_id"`
	CabangId            string `json:"cabang_id"`
	
}

type RiwayatProdukCabangRequest struct {
	Catatan string `json:"catatan"`
}

type ProdukCabangWithNote struct {
	Produk []ProdukCabangRequest `json:"produk"`
	Catatan RiwayatProdukCabangRequest `json:"catatan"`
}
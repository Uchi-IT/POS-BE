package dto

type ProdukGudangResponse struct {
	Id         string `json:"id"`
	Foto       string `json:"foto"`
	NamaProduk string `json:"nama_produk"`
	Stok       int    `json:"stok"`
	HargaJual  int    `json:"harga_jual"`
	HargaBeli  int    `json:"harga_beli"`
}

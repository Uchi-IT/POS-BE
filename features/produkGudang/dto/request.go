package dto

type ProdukGudangRequest struct {
	Foto       string `json:"foto"`
	NamaProduk string `json:"nama_produk"`
	Stok       int    `json:"stok"`
	HargaJual  int    `json:"harga_jual"`
	HargaBeli  int    `json:"harga_beli"`
}

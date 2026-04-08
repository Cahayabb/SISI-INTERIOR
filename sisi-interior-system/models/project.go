package models

type Project struct {
	ID               int    `json:"id"`
	JenisRuangan     string `json:"jenis_ruangan"`
	LuasRuangan      int    `json:"luas_ruangan"`
	TingkatKerumitan int    `json:"tingkat_kerumitan"`
	DurasiPengerjaan int    `json:"durasi_pengerjaan"`
	HargaProyek      int    `json:"harga_proyek"`
}

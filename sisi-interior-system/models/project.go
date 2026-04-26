package models

import "time"

type Project struct {
	ID                int       `json:"id" gorm:"primaryKey"`
	NamaPerusahaan    string    `json:"nama_perusahaan" gorm:"column:nama_perusahaan"`
	JenisProyek       string    `json:"jenis_proyek"`
	LuasArea          int       `json:"luas_area"`
	JenisPekerjaan    string    `json:"jenis_pekerjaan"`
	SpesifikasiDesign string    `json:"spesifikasi_design"`
	AreaLayanan       string    `json:"area_layanan"`
	DaftarItem        string    `json:"daftar_item"`
	SpesifikasiItem   string    `json:"spesifikasi_item"`
	HargaSatuan       int       `json:"harga_satuan"`
	Tanggal           time.Time `json:"tanggal"`
	CreatedAt         time.Time
}

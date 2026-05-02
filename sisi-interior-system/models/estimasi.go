package models

import "time"

type Estimasi struct {
	ID                int       `json:"id"`
	NamaPerusahaan    string    `json:"nama_perusahaan"`
	LuasArea          float64   `json:"luas_area"`
	JenisRuangan      string    `json:"jenis_ruangan"`
	JenisPekerjaan    string    `json:"jenis_pekerjaan"`
	TingkatKerumitan  string    `json:"tingkat_kerumitan"`
	DurasiPengerjaan  int       `json:"durasi_pengerjaan"`
	LokasiProyek      string    `json:"lokasi_proyek"`
	SpesifikasiDesign string    `json:"spesifikasi_design"`
	MaterialKhusus    string    `json:"material_khusus"`
	KebutuhanTambahan string    `json:"kebutuhan_tambahan"`
	HargaProyek       int64     `json:"harga_proyek"`
	CreatedAt         time.Time `json:"created_at"`
}

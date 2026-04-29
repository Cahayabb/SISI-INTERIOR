package models

import "time"

type Estimasi struct {
	ID                int       `json:"id"`
	NamaPerusahaan    string    `json:"nama_perusahaan"`
	LuasArea          float64   `json:"luas_area"`
	TingkatKerumitan  string    `json:"tingkat_kerumitan"`
	DurasiPengerjaan  int       `json:"durasi_pengerjaan"`
	JenisRuangan      string    `json:"jenis_ruangan"`
	JenisPekerjaan    string    `json:"jenis_pekerjaan"`
	SpesifikasiDesign string    `json:"spesifikasi_design"`
	LokasiProyek      string    `json:"lokasi_proyek"`
	MaterialKhusus    string    `json:"material_khusus"`
	DaftarItem        string    `json:"daftar_item"`
	KebutuhanTambahan string    `json:"kebutuhan_tambahan"`
	HargaProyek       float64   `json:"harga_proyek"`
	CreatedAt         time.Time `json:"created_at"`
}

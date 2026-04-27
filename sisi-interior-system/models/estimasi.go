package models

type Estimasi struct {
	ID                int     `json:"id"`
	NamaProyek        string  `json:"nama_proyek"`
	LuasArea          float64 `json:"luas_area"`
	TingkatKerumitan  int     `json:"tingkat_kerumitan"`
	DurasiPengerjaan  int     `json:"durasi_pengerjaan"`
	JenisRuangan      string  `json:"jenis_ruangan"`
	JenisPekerjaan    string  `json:"jenis_pekerjaan"`
	SpesifikasiDesign string  `json:"spesifikasi_design"`
	LokasiProyek      string  `json:"lokasi_proyek"`
	MaterialKhusus    string  `json:"material_khusus"`
	KebutuhanTambahan string  `json:"kebutuhan_tambahan"`
	HargaEstimasi     float64 `json:"harga_estimasi"`
	CreatedAt         string  `json:"created_at"`
}

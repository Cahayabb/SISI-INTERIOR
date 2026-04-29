package controllers

import (
	"net/http"
	"sisi-interior-system/config"
	"sisi-interior-system/models"

	"github.com/gin-gonic/gin"
)

// ============================
// REQUEST STRUCT
// ============================
type EstimasiRequest struct {
	NamaPerusahaan    string   `json:"nama_perusahaan"`
	LuasArea          float64  `json:"luas_area"`
	TingkatKerumitan  string   `json:"tingkat_kerumitan"`
	DurasiPengerjaan  int      `json:"durasi_pengerjaan"`
	JenisRuangan      string   `json:"jenis_ruangan"`
	JenisPekerjaan    string   `json:"jenis_pekerjaan"`
	SpesifikasiDesign string   `json:"spesifikasi_design"`
	LokasiProyek      string   `json:"lokasi_proyek"`
	MaterialKhusus    *string  `json:"material_khusus"`
	DaftarItem        []string `json:"daftar_item"`
	KebutuhanTambahan string   `json:"kebutuhan_tambahan"`
}

// ============================
// POST /estimasi
// ============================
func EstimasiBiaya(c *gin.Context) {
	var req EstimasiRequest

	// 1. Bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Validasi
	if req.NamaPerusahaan == "" || req.LuasArea <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nama perusahaan dan luas area wajib diisi",
		})
		return
	}

	// 3. Mapping kerumitan
	var nilaiKerumitan float64
	switch req.TingkatKerumitan {
	case "Mudah":
		nilaiKerumitan = 1
	case "Sedang":
		nilaiKerumitan = 1.5
	case "Sulit":
		nilaiKerumitan = 2
	default:
		nilaiKerumitan = 1
	}

	// 4. Mapping jenis pekerjaan
	var faktorPekerjaan float64
	switch req.JenisPekerjaan {
	case "Design":
		faktorPekerjaan = 0.5
	case "Build / Fit Out":
		faktorPekerjaan = 1.5
	case "Furniture":
		faktorPekerjaan = 1.2
	case "MEP (Mechanical Electrical Plumbing)":
		faktorPekerjaan = 1.3
	case "Maintenance":
		faktorPekerjaan = 0.8
	default:
		faktorPekerjaan = 1
	}

	// 5. Hitung jumlah item
	jumlahItem := float64(len(req.DaftarItem))

	// 6. Estimasi dasar
	basePrice := req.LuasArea * 1000000

	// 7. Final estimasi
	estimasi := (basePrice * nilaiKerumitan * faktorPekerjaan) +
		(jumlahItem * 500000) +
		(float64(req.DurasiPengerjaan) * 200000)

	// 8. Convert daftar item ke string (kalau DB masih string)
	daftarItemStr := ""
	if len(req.DaftarItem) > 0 {
		daftarItemStr = req.DaftarItem[0]
		for i := 1; i < len(req.DaftarItem); i++ {
			daftarItemStr += ", " + req.DaftarItem[i]
		}
	}

	// 9. Mapping ke model
	data := models.Estimasi{
		NamaPerusahaan:    req.NamaPerusahaan,
		LuasArea:          req.LuasArea,
		TingkatKerumitan:  req.TingkatKerumitan,
		DurasiPengerjaan:  req.DurasiPengerjaan,
		JenisRuangan:      req.JenisRuangan,
		JenisPekerjaan:    req.JenisPekerjaan,
		SpesifikasiDesign: req.SpesifikasiDesign,
		LokasiProyek:      req.LokasiProyek,
		KebutuhanTambahan: req.KebutuhanTambahan,
		HargaProyek:       estimasi,
		DaftarItem:        daftarItemStr,
	}

	// optional field
	if req.MaterialKhusus != nil {
		data.MaterialKhusus = *req.MaterialKhusus
	}

	// 10. Simpan ke DB
	result := config.DB.Table("data_proyek").Create(&data)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// 11. Response
	c.JSON(http.StatusOK, gin.H{
		"message":        "Data berhasil disimpan",
		"estimasi_harga": estimasi,
	})
}

// ============================
// GET /estimasi
// ============================
func GetEstimasi(c *gin.Context) {
	var data []models.Estimasi

	result := config.DB.Table("data_proyek").
		Order("id DESC").
		Find(&data)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

// ============================
// DELETE /estimasi
// ============================
func DeleteEstimasi(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID tidak ditemukan",
		})
		return
	}

	result := config.DB.Table("data_proyek").
		Where("id = ?", id).
		Delete(nil)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil dihapus",
	})
}

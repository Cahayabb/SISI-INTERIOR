package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	//"sisi-interior"
	"sisi-interior-system/config"
	"sisi-interior-system/models"
	"strings"

	"github.com/gin-gonic/gin"
)

// ============================
// REQUEST STRUCT
// ============================
type EstimasiRequest struct {
	NamaPerusahaan    string   `json:"nama_perusahaan"`
	LuasArea          float64  `json:"luas_area"`
	JenisRuangan      []string `json:"jenis_ruangan"`
	JenisPekerjaan    string   `json:"jenis_pekerjaan"`
	TingkatKerumitan  string   `json:"tingkat_kerumitan"`
	DurasiPengerjaan  int      `json:"durasi_pengerjaan"`
	LokasiProyek      string   `json:"lokasi_proyek"`
	SpesifikasiDesign string   `json:"spesifikasi_design"`
	MaterialKhusus    *string  `json:"material_khusus"`
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
	jumlahItem := float64(len(req.JenisRuangan))

	// 6. Estimasi dasar
	basePrice := req.LuasArea * 1000000

	// 7. Final estimasi
	estimasi := (basePrice * nilaiKerumitan * faktorPekerjaan) +
		(jumlahItem * 500000) +
		(float64(req.DurasiPengerjaan) * 200000)

	JenisRuanganStr := strings.Join(req.JenisRuangan, ", ")

	// 9. Mapping ke model
	data := models.Estimasi{
		NamaPerusahaan:    req.NamaPerusahaan,
		LuasArea:          req.LuasArea,
		JenisRuangan:      JenisRuanganStr,
		JenisPekerjaan:    req.JenisPekerjaan,
		TingkatKerumitan:  req.TingkatKerumitan,
		DurasiPengerjaan:  req.DurasiPengerjaan,
		LokasiProyek:      req.LokasiProyek,
		SpesifikasiDesign: req.SpesifikasiDesign,
		KebutuhanTambahan: req.KebutuhanTambahan,
		HargaProyek:       int64(estimasi),
	}

	// optional field
	if req.MaterialKhusus != nil {
		data.MaterialKhusus = *req.MaterialKhusus
	}

	// 10. Simpan ke DB
	result := config.DB.Table("estimasi_proyek").Create(&data)
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

	result := config.DB.Table("estimasi_proyek").
		Order("created_at DESC").
		Find(&data)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	var hasil []gin.H

	for _, item := range data {

		payload := map[string]interface{}{
			"luas_area":          item.LuasArea,
			"tingkat_kerumitan":  item.TingkatKerumitan,
			"durasi_pengerjaan":  item.DurasiPengerjaan,
			"jenis_ruangan":      item.JenisRuangan,
			"jenis_pekerjaan":    item.JenisPekerjaan,
			"spesifikasi_design": item.SpesifikasiDesign,
		}

		jsonData, _ := json.Marshal(payload)

		resp, err := http.Post(
			"http://127.0.0.1:5000/predict",
			"application/json",
			bytes.NewBuffer(jsonData),
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)

		var mlResult map[string]interface{}
		json.Unmarshal(body, &mlResult)

		hasil = append(hasil, gin.H{
			"data":     item,
			"estimasi": mlResult["estimasi"],
		})
	}

	c.JSON(http.StatusOK, hasil)
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

	result := config.DB.Table("estimasi_proyek").
		Where("id = ?", id).
		Delete(&models.Estimasi{})

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

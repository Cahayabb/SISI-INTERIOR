package controllers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

var DB *sql.DB

// ============================
// REQUEST STRUCT
// ============================
type EstimasiRequest struct {
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
}

// ============================
// POST /estimasi
// ============================
func EstimasiBiaya(c *gin.Context) {
	var req EstimasiRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	// ============================
	// CALL PYTHON API
	// ============================
	jsonData, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal encode data"})
		return
	}

	resp, err := http.Post("http://localhost:5000/predict", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ML service error"})
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal decode response ML"})
		return
	}

	estimasiVal, ok := result["estimasi"].(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Format response ML tidak valid"})
		return
	}

	// ============================
	// INSERT DATABASE
	// ============================
	query := `
	INSERT INTO data_proyek 
	(nama_proyek, luas_area, tingkat_kerumitan, durasi_pengerjaan,
	 jenis_ruangan, jenis_pekerjaan, spesifikasi_design,
	 lokasi_proyek, material_khusus, kebutuhan_tambahan, harga_proyek)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
	`

	_, err = DB.Exec(query,
		req.NamaProyek,
		req.LuasArea,
		req.TingkatKerumitan,
		req.DurasiPengerjaan,
		req.JenisRuangan,
		req.JenisPekerjaan,
		req.SpesifikasiDesign,
		req.LokasiProyek,
		req.MaterialKhusus,
		req.KebutuhanTambahan,
		estimasiVal,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal simpan data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"estimasi_harga": estimasiVal,
	})
}

// ============================
// GET /estimasi (Dashboard)
// ============================
func GetEstimasi(c *gin.Context) {
	rows, err := DB.Query("SELECT id, nama_proyek, harga_proyek, created_at FROM data_proyek ORDER BY id DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal ambil data"})
		return
	}
	defer rows.Close()

	var list []map[string]interface{}

	for rows.Next() {
		var id int
		var nama string
		var harga float64
		var tanggal string

		if err := rows.Scan(&id, &nama, &harga, &tanggal); err != nil {
			continue
		}

		list = append(list, gin.H{
			"id":          id,
			"nama_proyek": nama,
			"harga":       harga,
			"tanggal":     tanggal,
		})
	}

	c.JSON(http.StatusOK, list)
}

// ============================
// DELETE /estimasi (Dashboard)
// ============================
func DeleteEstimasi(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID tidak ditemukan",
		})
		return
	}

	_, err := DB.Exec("DELETE FROM data_proyek WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(), // tampilkan error asli biar kelihatan
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil dihapus",
	})
}

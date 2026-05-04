package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"sisi-interior-system/models"

	"github.com/gin-gonic/gin"
)

// ============================
// ESTIMASI ONLY USERS
// TANPA SIMPAN DB
// ============================
func EstimasiOnly(c *gin.Context) {
	var req models.Estimasi

	// 1. bind input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 2. kirim ke ML (Python API)
	jsonData, _ := json.Marshal(req)

	resp, err := http.Post(
		"http://127.0.0.1:5000/predict",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal menghubungi ML service",
		})
		return
	}
	defer resp.Body.Close()

	// 3. decode response ML
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	// 4. return ONLY hasil (NO DB)
	c.JSON(http.StatusOK, gin.H{
		"message":        "Estimasi berhasil",
		"estimasi_harga": result["estimasi"],
	})
}

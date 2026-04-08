package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type EstimasiRequest struct {
	LuasArea         int `json:"luas_area"`
	TingkatKerumitan int `json:"tingkat_kerumitan"`
	DurasiPengerjaan int `json:"durasi_pengerjaan"`
}

func EstimasiBiaya(c *gin.Context) {

	var req EstimasiRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Input tidak valid",
		})
		return
	}

	// sementara pakai rumus sederhana
	estimasi := (req.LuasArea * 500000) +
		(req.TingkatKerumitan * 2000000) +
		(req.DurasiPengerjaan * 300000)

	c.JSON(http.StatusOK, gin.H{
		"estimasi_harga": estimasi,
	})
}

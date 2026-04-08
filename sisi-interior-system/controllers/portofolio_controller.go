package controllers

import (
	"net/http"
	"sisi-interior-system/config"
	"sisi-interior-system/models"

	"github.com/gin-gonic/gin"
)

func CreatePortfolio(c *gin.Context) {

	var portfolio models.Portfolio

	if err := c.ShouldBindJSON(&portfolio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	query := `
	INSERT INTO portfolio (judul_proyek, deskripsi, gambar)
	VALUES ($1,$2,$3)
	`

	_, err := config.DB.Exec(
		query,
		portfolio.JudulProyek,
		portfolio.Deskripsi,
		portfolio.Gambar,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal menyimpan portfolio",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Portfolio berhasil ditambahkan",
	})
}

func GetPortfolio(c *gin.Context) {

	rows, err := config.DB.Query(`
	SELECT id, judul_proyek, deskripsi, gambar
	FROM portfolio
	ORDER BY id DESC
	`)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengambil data portfolio",
		})
		return
	}

	defer rows.Close()

	var portfolios []models.Portfolio

	for rows.Next() {
		var p models.Portfolio

		err := rows.Scan(
			&p.ID,
			&p.JudulProyek,
			&p.Deskripsi,
			&p.Gambar,
		)

		if err != nil {
			continue
		}

		portfolios = append(portfolios, p)
	}

	c.JSON(http.StatusOK, portfolios)
}

func UpdatePortfolio(c *gin.Context) {

	id := c.Param("id")

	var portfolio models.Portfolio

	if err := c.ShouldBindJSON(&portfolio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	query := `
	UPDATE portfolio
	SET judul_proyek=$1, deskripsi=$2, gambar=$3
	WHERE id=$4
	`

	_, err := config.DB.Exec(
		query,
		portfolio.JudulProyek,
		portfolio.Deskripsi,
		portfolio.Gambar,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal update portfolio",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Portfolio berhasil diupdate",
	})
}

func DeletePortfolio(c *gin.Context) {

	id := c.Param("id")

	query := `DELETE FROM portfolio WHERE id=$1`

	_, err := config.DB.Exec(query, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal menghapus portfolio",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Portfolio berhasil dihapus",
	})
}

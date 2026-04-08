package models

type Portfolio struct {
	ID          int    `json:"id"`
	JudulProyek string `json:"judul_proyek"`
	Deskripsi   string `json:"deskripsi"`
	Gambar      string `json:"gambar"`
}

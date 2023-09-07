package models

type Playlist struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Song   string `json:"song"`
	Artist string `json:"artist"`
	Info   string `json:"info"`
	URL    string `json:"url"`
}

type CreatePlaylist struct {
	Song   string `json:"song" binding:"required"`
	Artist string `json:"artist" binding:"required"`
	Info   string `json:"info" binding:"required"`
	URL    string `json:"url" binding:"required"`
}

type UpdatePlaylist struct {
	Song   string `json:"song"`
	Artist string `json:"artist"`
	Info   string `json:"info"`
	URL    string `json:"url"`
}

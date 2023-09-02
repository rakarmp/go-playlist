package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type playlist struct {
	ID     string `json:"id"`
	Song   string `json:"song"`
	Artist string `json:"artist"`
	URL    string `json:"url"`
}

func getPlaylist(c *gin.Context) {
	var playlist = []playlist{
		{ID: "1", Song: "Wind", Artist: "明星 Akeboshi", URL: "https://open.spotify.com/track/5BqKtuCFLfZyzfZOwlgW1f"},
		{ID: "2", Song: "Dandelions", Artist: "Ruth B.", URL: "https://open.spotify.com/track/2eAvDnpXP5W0cVtiI0PUxV"},
		{ID: "3", Song: "Atlantis", Artist: "Seafret", URL: "https://open.spotify.com/track/1Fid2jjqsHViMX6xNH70hE"},
	}

	c.IndentedJSON(http.StatusOK, playlist)
}

func getWelcome(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
	})
}

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", getWelcome)
	router.GET("/playlist", getPlaylist)

	router.Run("localhost:8000")
}

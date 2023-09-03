package main

import (
	"example/first-web-service/controllers"
	"example/first-web-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Base Gin
	router := gin.Default()
	// Models
	models.ConnectDatabase()

	// Router
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Simple Playlist API, Pure GO",
		})
	})
	router.GET("/playlist", controllers.FindPlaylist)
	router.POST("/playlist", controllers.CreatePlaylist)
	router.GET("/playlist/:id", controllers.FindPlaylistId)
	router.PATCH("playlist/:id", controllers.UpdatePlaylist)
	router.DELETE("/playlist/:id", controllers.DeletePlaylist)

	// Proxy
	router.SetTrustedProxies(nil)

	// Server
	router.Run("localhost:8000")
}

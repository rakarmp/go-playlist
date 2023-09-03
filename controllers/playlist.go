package controllers

import (
	"example/first-web-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /playlist

func FindPlaylist(ctx *gin.Context) {
	var playlist []models.Playlist
	models.DB.Find(&playlist)

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": playlist,
	})
}

func FindPlaylistId(ctx *gin.Context) {
	var playlist models.Playlist

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&playlist).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Not Found!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": playlist,
	})
}

// POST /playlist

func CreatePlaylist(ctx *gin.Context) {
	var input models.CreatePlaylist
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	playlist := models.Playlist{
		Song:   input.Song,
		Artist: input.Artist,
		URL:    input.URL,
	}
	models.DB.Create(&playlist)

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": playlist,
	})
}

// Update

func UpdatePlaylist(ctx *gin.Context) {
	var playlist models.Playlist

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&playlist).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Not Found!",
		})
		return
	}

	var input models.UpdatePlaylist
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	models.DB.Model(&playlist).Updates(input)

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": playlist,
	})
}

// Delete

func DeletePlaylist(ctx *gin.Context) {
	var playlist models.Playlist

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&playlist).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Not Found!",
		})
		return
	}

	models.DB.Delete(&playlist)

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": true,
	})
}

package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ybuilds/slash/models"
	"github.com/ybuilds/slash/utils"
)

func CreateMapping(ctx *gin.Context) {
	var url models.Url

	err := ctx.ShouldBindJSON(&url)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	id, err := url.CreateMapping()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"url": id})
}

func GetMapping(ctx *gin.Context) {
	encode := ctx.Param("encode")
	if encode == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid url, no encode found"})
		return
	}

	id := utils.Base62Decoder(encode)
	url, err := models.GetMapping(id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"url": url})
}

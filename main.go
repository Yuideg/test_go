package main

import (
	"github.com/Yideg/test_go/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.GET("/albums", postAlbums)
	port := os.Getenv("PORT")
	router.Run(":" + port)
}

//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
func getAlbums(c *gin.Context) {
	new := &models.Test{
		Message: "Get all Albums",
	}
	c.JSON(http.StatusOK, new)
}
func getAlbumByID(c *gin.Context) {
	new := &models.Test{
		Message: "Get Album By ID",
	}
	c.JSON(http.StatusOK, new)
}
func postAlbums(c *gin.Context) {
	new := &models.Test{
		Message: " Posting new Albums",
	}
	c.JSON(http.StatusOK, new)
}

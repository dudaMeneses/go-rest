package main

import (
	"github.com/gin-gonic/gin"
	"go-rest/handler"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handler.NewHandler().GetAlbums)

	router.Run("localhost:8080")
}
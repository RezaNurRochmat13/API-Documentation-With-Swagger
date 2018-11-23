package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func main() {
	router := gin.New()

	router.GET("users", Users)
	// Use ginSwagger middleware to
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}

func Users(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Data users"})
}

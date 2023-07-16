package main

import (
	"summarize/controller"
	"summarize/model"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	model.ConnectDatabase()

	router.POST("/posts", controller.CreateText)
	router.Run("localhost:8080")
}

package controller

import (
	"net/http"
	"summarize/model"
	"summarize/service"

	"github.com/gin-gonic/gin"
)

type CreateTextInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func CreateText(c *gin.Context) {
	var input CreateTextInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sum_response := service.SummarizeText(input.Content)
	post := model.Text{Title: input.Title, Content: input.Content, Summary: sum_response.Summary}
	model.DB.Create(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

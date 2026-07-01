package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (p PongController) Pong(c *gin.Context) {
	fmt.Println("My handler")
	c.JSON(http.StatusOK, gin.H{
		"message": "ping...pong",
	})
}

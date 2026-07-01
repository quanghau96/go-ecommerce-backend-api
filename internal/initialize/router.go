package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin"
	c "github.com/quanghau96/go-ecommerce-backend-api/internal/controller"
	"github.com/quanghau96/go-ecommerce-backend-api/internal/middlewares"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.AuthenMiddleware(), BB(), CC)

	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", c.NewPongController().Pong)
		v1.GET("/user/1", c.NewUserController().GetUserById)
		// v1.PUT("/ping", Pong)
		// v1.PATCH("/ping", Pong)
		// v1.DELETE("/ping", Pong)
		// v1.HEAD("/ping", Pong)
		// v1.OPTIONS("/ping", Pong)
	}

	return r
}

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before AA")
		c.Next()
		fmt.Println("After AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before BB")
		c.Next()
		fmt.Println("After BB")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before CC")
	c.Next()
	fmt.Println("After CC")
}

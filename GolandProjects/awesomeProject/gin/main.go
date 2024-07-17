package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})

	server.GET("/hello/:name", func(ctx *gin.Context) {

		ctx.String(http.StatusOK, "hello "+ctx.Param("name"))
	})

	server.GET("/order", func(ctx *gin.Context) {
		id := ctx.Query("id")
		ctx.String(http.StatusOK, "订单 ID 是"+id)
	})

	server.GET("/views/*.html", func(ctx *gin.Context) {
		view := ctx.Param(".html")
		ctx.String(http.StatusOK, "views 是 "+view)
	})

	server.POST("/login", func(ctx *gin.Context) {

		ctx.String(http.StatusOK, "hello,login")
	})

	server.Run(":8080")

}

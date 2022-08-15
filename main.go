package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	handler := gin.New()
	handler.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]any{
			"status": "OK",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":"+port, handler)

}

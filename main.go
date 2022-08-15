package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	handler := gin.Default()
	handler.GET("/", indexPage)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":"+port, handler)

}

func indexPage(ctx *gin.Context) {
	ctx.JSON(200, map[string]any{
		"status": "OK",
	})
}

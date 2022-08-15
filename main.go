package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	handler := gin.Default()
	handler.GET("/", indexPage)
	handler.GET("/audio/stream", streamAudio)

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

func indexstreamAudioPage(ctx *gin.Context) {
	f, err := os.Open("audio.mp3")
	if err != nil {
		ctx.Status(500)
		return
	}
	defer f.Close()

	ctx.Stream(func(w io.Writer) bool {
		_, err := io.Copy(w, f)
		return err == nil
	})
}

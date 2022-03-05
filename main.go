package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/pod-name", PodName)

	err := r.Run(":8080")
	if err != nil {
		log.Printf("run: %s\n", err)
	}
}

func PodName(ctx *gin.Context) {
	pn := os.Getenv("POD_NAME")
	if pn == "" {
		pn = "pod name not found"
	}
	ctx.JSON(http.StatusOK, gin.H{
		"pod_name": pn,
	})
}

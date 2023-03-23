package controller

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetAllResources(ctx *gin.Context) {
	file, err := os.OpenFile("data/resources.json", os.O_RDONLY, os.ModePerm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer file.Close()
	io.Copy(ctx.Writer, file)
	ctx.Status(http.StatusOK)
}

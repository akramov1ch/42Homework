package main

import (
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/server2?userid=25", func(contex *gin.Context) {
		responce, err := http.Get("localhost:8080")
		if err != nil {
			contex.JSON(400, err)
		}

		data, err := io.ReadAll(responce.Body)
		if err != nil {
			contex.JSON(400, err)
		}
		contex.IndentedJSON(200, gin.H{
			"message": string(data),
		})
	})

	r.Run(":9000")
}

package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()
	route.GET("/", func(contex *gin.Context) {
		id := contex.Param("id")
		user_id := contex.Query("user_id")
		contex.JSON(200, gin.H{
			"id":      id,
			"user_id": user_id,
		})
	})
	route.Run(":8080")
}

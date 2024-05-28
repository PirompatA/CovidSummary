package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/covid/summary", func(ctx *gin.Context) {

	})

	router.Run(":8080")
}

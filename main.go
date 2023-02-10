package main

import (
	"ShortUrl/app/shortUrl/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	server := gin.Default()
	server.GET("/:path", router.Redirect)
	server.POST("/add", router.GenerateUrl)
	server.POST("/analysis", router.AnalysisUrl)
	server.DELETE("/delete", router.DeleteUrl)
	err := server.Run(":8088")
	if err != nil {
		log.Fatalln("ListenAndServer error: ", err)
	}
}

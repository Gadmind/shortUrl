package main

import (
	"ShortUrl/app/shortUrl/router"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	server := gin.Default()
	server.Static("/static", "./static")
	server.LoadHTMLGlob("templates/**/*")
	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/index.html", gin.H{
			"title": "index/index",
		})
	})
	server.GET("/:path", router.Redirect)
	server.POST("/add", router.GenerateUrl)
	server.POST("/analysis", router.AnalysisUrl)
	server.DELETE("/delete", router.DeleteUrl)

	err := server.Run(":8088")
	if err != nil {
		log.Fatalln("ListenAndServer error: ", err)
	}
}

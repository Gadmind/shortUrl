package route

import (
	"ShortUrl/app/short_url"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func API() {
	serverDefault := gin.Default()
	serverDefault.Static("/static", "./static")
	serverDefault.LoadHTMLGlob("templates/**/*")
	serverDefault.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/index.html", gin.H{
			"title": "index/index",
		})
	})
	serverDefault.GET("/:path", short_url.Redirect)
	serverDefault.POST("/url/add", short_url.GenerateUrl)
	serverDefault.POST("/url/analysis", short_url.AnalysisUrl)
	serverDefault.POST("/url/delete", short_url.DeleteUrl)
	serverDefault.GET("/url/list", short_url.GetList)
	serverDefault.GET("/url/Delete", short_url.Delete)
	serverDefault.GET("/url/rank", short_url.Delete)

	server := &http.Server{
		Addr:    ":8088",
		Handler: serverDefault,
	}

	go func() {
		// 服务连接
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

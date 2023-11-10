package main

import (
	"github.com/AbirAhsan/gin-crash-course/controller"
	middlewares "github.com/AbirAhsan/gin-crash-course/middleware"
	"github.com/AbirAhsan/gin-crash-course/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	middlewares.SetUpLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	// server.GET("/posts", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "OK!",
	// 	})
	// })
	server.GET("/getVideos", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"success":   true,
			"videoList": VideoController.FindAll(),
		})

	})
	server.POST("/saveVideos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))
	})

	server.Run(":8080")
}

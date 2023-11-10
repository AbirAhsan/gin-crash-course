package controller

import (
	"github.com/AbirAhsan/gin-crash-course/entity"
	"github.com/AbirAhsan/gin-crash-course/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video

	Save(ctx *gin.Context) entity.Video
}
type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	videos := c.service.FindAll()

	if videos != nil {
		return videos
	}
	return []entity.Video{}
}
func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}

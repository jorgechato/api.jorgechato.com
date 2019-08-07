package api

import (
	"github.com/gin-gonic/gin"

	"jorgechato.com/api/location"
	"jorgechato.com/api/project"
	"jorgechato.com/api/subscription"
	. "jorgechato.com/utils"
)

func Build() (router *gin.Engine) {
	router = gin.Default()

	v := router.Group(APIBASE)
	{
		v.GET("/where", location.Today)
		v.GET("/projects", project.Repos)
		v.POST("/subscribe", subscription.List)
	}

	return
}

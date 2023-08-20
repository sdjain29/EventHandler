package routes

import (
	userhandler "EventHandler/core/eventIngest"

	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine) {

	authorised := r.Group("/")
	// authorised.Use(utils.BasicAuthUser)
	{
		authorised.POST("/event", userhandler.EventIngest)

	}

}

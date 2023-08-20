package routes

import (
	"EventHandler/utils"

	"github.com/gin-gonic/gin"
)

func Internal(r *gin.Engine) {
	authorisedInternalServer := r.Group("/internal")
	authorisedInternalServer.Use(utils.BasicAuthInternal)
	{
	}
}

package utils

import (
	"EventHandler/config"

	"github.com/gin-gonic/gin"
)

func BasicAuthUser(c *gin.Context) {
}

func BasicAuthInternal(c *gin.Context) {
	BasicAuth(c, "X-Internal-Token", config.Config.InternalServerToken)
}

func BasicAuth(c *gin.Context, s string, pass string) {
}

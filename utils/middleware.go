package utils

import (
	"github.com/gin-gonic/gin"
)

func AddRequestId(c *gin.Context) {
	defer recover()
	if c.GetHeader("X-Request-Id") != "" {
		c.Set("requestid", c.GetHeader("X-Request-Id"))
		c.Header("X-Request-Id", c.Value("requestid").(string))
	} else {
		c.Set("requestid", UuidGenerator())
		c.Header("X-Request-Id", c.Value("requestid").(string))
	}
	c.Next()
}

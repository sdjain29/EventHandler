package utils

import (
	"EventHandler/constants"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Marshal(c *gin.Context, s interface{}) string {
	out, err := json.Marshal(s)
	CheckErrAndCrash(c, constants.InternalServerErrorCode, err, fmt.Sprint(err))
	return string(out)
}

func MarshalWihtoutError(s interface{}) (string, error) {
	out, err := json.Marshal(s)
	return string(out), err
}

package utils

import (
	"EventHandler/constants"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func UnmarshalWithValue[T any](c *gin.Context, body []byte) T {
	var t T
	if len(body) > 0 {
		err := json.Unmarshal(body, &t)
		CheckErrAndCrash(c, constants.InternalServerErrorCode, err, fmt.Sprint(err))
	}
	return t
}

func Unmarshal(c *gin.Context, body []byte, s interface{}) {
	err := json.Unmarshal(body, &s)
	CheckErrAndCrash(c, constants.InternalServerErrorCode, err, fmt.Sprint(err))
}

func UnmarshalWithError(c *gin.Context, body []byte, s interface{}) error {
	err := json.Unmarshal(body, &s)
	return err
}

func Marshal(c *gin.Context, s interface{}) string {
	out, err := json.Marshal(s)
	CheckErrAndCrash(c, constants.InternalServerErrorCode, err, fmt.Sprint(err))
	return string(out)
}

func MarshalWihtoutError(s interface{}) (string, error) {
	out, err := json.Marshal(s)
	return string(out), err
}

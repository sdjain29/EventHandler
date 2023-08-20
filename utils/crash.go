package utils

import (
	"EventHandler/constants"
	core "EventHandler/init"
	"EventHandler/logger"
	"fmt"
	"reflect"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

type ErrorMsg struct {
	StatusCode int    `json:"statusCode"`
	ErrorMsg   string `json:"errorString"`
	UserMsg    string `json:"userMsg"`
}

func ThrowAndAbort(c *gin.Context, abortStatus int, err error, errorstring string) {
	var output ErrorMsg
	output.StatusCode = abortStatus
	output.ErrorMsg = fmt.Sprint(err)
	output.UserMsg = errorstring
	logger.Error(c, string(debug.Stack()))
	if fmt.Sprint(reflect.TypeOf(c.Writer)) == "*logger.ResponseBodyWriter" {
		c.AbortWithStatusJSON(abortStatus, output)
	}
	panic(nil)
}

func CheckErrAndCrash(c *gin.Context, abortStatus int, err error, errorstring string) {
	if err != nil {
		ThrowAndAbort(c, abortStatus, err, errorstring)
	}
}

func InputValidation(c *gin.Context, s interface{}) {
	err := core.Validate.Struct(s)
	CheckErrAndCrash(c, constants.ClientBadRequestCode, err, constants.BadRequest)
}

func ThreadRecovery() {
	recover()
}

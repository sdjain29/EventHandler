package eventIngest

import (
	"EventHandler/config"
	"EventHandler/storage"
	types "EventHandler/types/api/eventIngest"
	"EventHandler/utils"
	"time"

	//"crypto/rand"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func EventIngest(c *gin.Context) {
	var Input types.EventIngestRequest
	c.ShouldBindBodyWith(&Input, binding.JSON)
	eventIngestInputValidation(c, Input)
	eventIngestUtils(c, Input)
}

func eventIngestInputValidation(c *gin.Context, Input types.EventIngestRequest) {
	utils.InputValidation(c, Input)
	// for _, j := range Input.Destination {

	// 	if !utils.Some(config.Config.DestinationServer, equalityFn) {
	// 		utils.ThrowAndAbort(c, constants.ClientBadRequestCode, errors.New(constants.BadRequest), constants.BadRequest)
	// 	}
	// }
}

func eventIngestUtils(c *gin.Context, Input types.EventIngestRequest) {
	for _, j := range Input.Destination {
		equalityFn := func(s string) bool { return s == j }
		if !utils.Some(config.Config.DestinationServer, equalityFn) {
			var temp storage.Event
			temp.Id = utils.UuidGenerator()
			temp.CreatedAt = time.Now()
			temp.UpdatedAt = time.Now()
			temp.UserId = Input.UserId
			temp.Destination = j
			temp.Payload = Input.Payload
			temp.RetryCount = 0
			temp.EventTime = Input.EventTime
			temp.Status = "ACTIVE"
			utils.RedisSet(temp.Id, utils.Marshal(c, temp), time.Duration(24*time.Hour))
			utils.RedisSetZAdd(c, j, float64(time.Now().Unix()), temp.Id)
		}
	}
	var Output types.EventIngestResponse
	Output.Status = "SUCCESS"
	c.JSON(200, Output)
}

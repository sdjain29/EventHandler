package eventDelivery

import (
	"EventHandler/config"
	"EventHandler/storage"
	"EventHandler/utils"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type DestinationResponse struct {
	Status           string  `json:"status"`
	NextScheduleTime float64 `json:"nextScheduleTime"`
}

type DestEvent struct {
	UserId    string    `json:"userId"`
	Payload   string    `json:"payload"`
	EventTime time.Time `json:"eventTime"`
}

const MaxRetry int = 10

func ProcessDelivery() {
	for _, j := range config.Config.DestinationServer {
		go processDelivery(j)
	}
}

func processDelivery(schema string) {
	for {
		fmt.Println("looperrr ", schema)
		process := utils.RedisZGet(schema, fmt.Sprint(time.Now().Unix()))
		if len(process) == 0 {
			time.Sleep(20 * time.Second)
		}

		for _, j := range process {
			utils.RedisSetZDel(schema, j)
			event, err := utils.RedisGet(j)
			fmt.Println(event)
			if err == nil && event != "" {
				var temp storage.Event
				err := json.Unmarshal([]byte(event), &temp)
				if err == nil {
					switch temp.Status {
					case "COMPLETED":
					default:
					case "ACTIVE":
						temp.Status = "QUEUED"
						temp.RetryCount = temp.RetryCount + 1
						updateEvent(temp)
						utils.RedisSetZAddWithoutContext(schema, float64(time.Now().Add(6*time.Hour).Unix()), temp.Id)
						pushEvent(temp)
					case "QUEUED":
						if time.Now().Sub(temp.UpdatedAt).Hours() > 6 {
							temp.RetryCount = temp.RetryCount + 1
							updateEvent(temp)
							utils.RedisSetZAddWithoutContext(schema, float64(time.Now().Add(6*time.Hour).Unix()), temp.Id)
							pushEvent(temp)
						}
					}
				}
			}
		}
	}
}

func updateEvent(event storage.Event) {
	event.UpdatedAt = time.Now()
	out, err := utils.MarshalWihtoutError(event)
	if err != nil {
		return
	}
	utils.RedisSet(event.Id, out, time.Duration(24*time.Hour))
}

func pushEvent(event storage.Event) {
	if event.RetryCount > MaxRetry {
		utils.RedisSetZDel(event.Destination, event.Id)
		return
	}
	switch event.Destination {
	case "DEST1", "DEST2", "DEST3":
		output := dummypushEventUtils(DestEvent{UserId: event.UserId, EventTime: event.EventTime}, event.RetryCount)
		utils.RedisSetZDel(event.Destination, event.Id)
		utils.RedisSetZAddWithoutContext(event.Destination, output.NextScheduleTime, event.Id)
		event.Status = output.Status
		updateEvent(event)
	}
}

func dummypushEventUtils(event DestEvent, retryCount int) DestinationResponse {
	var output DestinationResponse
	output.Status = "ACTIVE"
	output.NextScheduleTime = float64(time.Now().Add(time.Minute * time.Duration(retryCount*5)).Unix())

	prob := rand.Intn(100) + 1

	switch {
	case prob < 50: // success scenarios
		output.Status = "COMPLETED"
	case prob < 75: // failure scenarios
		output.NextScheduleTime = float64(time.Now().Add(time.Minute * 5).Unix())
	default: // delay, etc scenarios
		time.Sleep(1 * time.Minute)
		// base
	}

	return output
}

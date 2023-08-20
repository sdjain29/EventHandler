package eventIngest

import "time"

type EventIngestRequest struct {
	UserId      string    `json:"userId" validate:"required"`
	Payload     string    `json:"payload" validate:"required"`
	Destination []string  `json:"destination" validate:"required"`
	EventTime   time.Time `json:"eventTime" validate:"required"`
}

type EventIngestResponse struct {
	Status string `json:"status"`
}

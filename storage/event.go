package storage

import (
	"time"
)

type Event struct {
	Id          string    `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UserId      string    `json:"userId"`
	Payload     string    `json:"payload"`
	Destination string    `json:"destination"`
	RetryCount  int       `json:"retryCount"`
	EventTime   time.Time `json:"eventTime"`
	Status      string    `json:"status"`
}

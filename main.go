package main

import (
	"EventHandler/config"
	processhandler "EventHandler/core/eventDelivery"
	core "EventHandler/init"
	"EventHandler/logger"
	route "EventHandler/routes"
	"EventHandler/utils"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {

	config.SetConfig()
	utils.SetSeedForRand()
	core.SetupValidation()
	logger.SetupLogger()

	var redisClient *redis.Client
	redisClient = core.SetupRedis()

	defer func() {
		if redisClient != nil {
			redisClient.Close()
		}
	}()

	switch os.Getenv("MODE") {
	case "EVENTHANDLER":
		go processhandler.ProcessDelivery()
		EventHandler()
	// case "EventHandler/KAFKAPRODUCER":
	// 	processhandler.ProcessScheduler()
	// case "EventHandler/KAFKACONSUMER":
	// 	go processhandler.ConsumerLeasingLock()
	// 	go processhandler.ConsumerAccountTransfer()
	// 	processhandler.ConsumerGroup()
	default:
		log.Println(os.Getenv("MODE"))
	}
}

func EventHandler() {
	defer utils.ThreadRecovery()
	r := gin.Default()
	s := &http.Server{
		Addr:         ":11111",
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	r.Use(utils.AddRequestId)
	r.Use(logger.GinLogger)
	r.Use(gin.Recovery())

	route.User(r)
	route.Internal(r)
	s.ListenAndServe()
}

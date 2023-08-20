package config

func dev() configType {
	var config configType
	config.Redis.RedisHost = []string{"host.docker.internal:6379"}
	config.Redis.RedisPassword = ""
	config.Redis.RedisDb = 0
	config.DestinationServer = []string{"DEST1", "DEST2", "DEST3"}
	return config
}

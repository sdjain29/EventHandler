package config

type configType struct {
	Redis               RedisConfig
	InternalServerToken string
	DestinationServer   []string
}

type RedisConfig struct {
	RedisHost     []string
	RedisPassword string
	RedisDb       int
}

var Config configType

func SetConfig() {
	Config = dev()
}

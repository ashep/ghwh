package server

type Config struct {
	Address     string `envconfig:"SERVER_ADDRESS" default:":9000"`
	ReadTimeout int    `envconfig:"SERVER_READ_TIMEOUT" default:"5"`
}

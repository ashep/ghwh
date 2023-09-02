package server

type Config struct {
	Address     string `default:":9000"`
	ReadTimeout int    `default:"5"`
	AuthToken   string
	Debug       bool
}
